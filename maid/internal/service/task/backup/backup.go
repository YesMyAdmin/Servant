package backup

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	backupEntity "common/public/model/entity/backup"
	backupsvc "common/public/service/backup"
	"common/public/utils"
	"maid/internal/service/task"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"golang.org/x/crypto/ssh"
)


// DoBackup 执行备份任务
func DoBackup(taskId uint64) error {
	//加载备份任务信息
	backupTask, err := backupsvc.GetBackupTask(taskId)
	if (err != nil) {
		return err
	}
	//获取备份存储记录,构造备份存储器列表
	//一个备份任务会下挂多个备份存储，需要将同一文件依次上传
	dumps, listDumpsErr := backupsvc.ListDumpsByTaskId(taskId)
	if (listDumpsErr != nil) {
		return listDumpsErr
	}
	dumpers := make([]BackupDumper, 0, len(dumps))
	for _, dump := range dumps {
		dumpers = append(dumpers, makeDumper(dump))
	}

	if (backupTask.Mode == backupEntity.Full) {
		err := fullBackup(backupTask, dumpers)
		if (err != nil) {
			return err
		}
	} else {
		//增量备份模式
	}
	return nil
}

// fullBackup 全量备份
func fullBackup(backupTask *backupEntity.BackupTask, dumpers []BackupDumper) error {
	taskPool, err := task.NewTaskPool()
	if (err != nil) {
		return err
	}
	//生成备份记录id
	backupId := 0
	//全量备份模式
	file, err := os.Open(backupTask.Source)
	if (err != nil) {
		return err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if (err != nil) {
		return err
	}

	var sourcePath string
	fileName := filepath.Base(file.Name()) + "_" + strconv.Itoa(backupId) + ".zip"
	if (fileInfo.IsDir()) {
		//如果目标是文件夹,将文件压缩后上传
		temporaryZipFile := "./cached/backup/" + fileName
		err := utils.ZipFolder(backupTask.Source, temporaryZipFile)
		if (err != nil) {
			return err
		}
		sourcePath = temporaryZipFile
	} else {
		//如果目标是文件,直接上传
		sourcePath = backupTask.Source
	}

	// 依次将同一文件转储到每个存储位置,单个转储失败不影响其他转储
	var wg sync.WaitGroup
	var mu sync.Mutex
	var dumpErrs []error
	for _, dumper := range dumpers {
		wg.Add(1)
		taskPool.Submit(func() {
			defer wg.Done()
			if _, dumpErr := dumper.Dump(sourcePath, fileName); dumpErr != nil {
				slog.Error("备份文件转储失败", slog.Any("error", dumpErr))
				mu.Lock()
				dumpErrs = append(dumpErrs, dumpErr)
				mu.Unlock()
			}
		})
	}
	wg.Wait()
	if len(dumpErrs) > 0 {
		return errors.Join(dumpErrs...)
	}
	return nil
}

// 获取一个转储器实例
func makeDumper(dumpMetadata *backupEntity.BackupDump) (BackupDumper) {
	switch dumpMetadata.DumpType {
	case backupEntity.S3:
		return &S3Dumper{dumpMetadata: dumpMetadata}
	case backupEntity.SCP:
		return &ScpDumper{dumpMetadata: dumpMetadata}
	default:
		return &LocalDumper{dumpMetadata: dumpMetadata}
	}
}

// 备份文件转储抽象接口
type BackupDumper interface {
	// Dump 将 sourcePath 指向的文件转储到目标位置,uploadedName 为转储后的文件名
	Dump(sourcePath, uploadedName string) (*backupEntity.BackupFileRecord, error)
}

// S3协议转储器
type S3Dumper struct {
	dumpMetadata *backupEntity.BackupDump
}

// 转储文件
func (s *S3Dumper) Dump(sourcePath, uploadedName string) (*backupEntity.BackupFileRecord, error) {
	if s.dumpMetadata.Credential == nil {
		return nil, errors.New("s3转储认证信息缺失")
	}
	credential, ok := (*s.dumpMetadata.Credential).(*backupEntity.S3DumpCredential)
	if !ok {
		return nil, errors.New("s3转储认证信息无效")
	}
	endpoint, bucket, prefix, err := parseS3Target(s.dumpMetadata.UrlTemplate)
	if (err != nil) {
		return nil, err
	}
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(credential.AccessKeyId, credential.AccessKey, ""),
		Secure: strings.HasPrefix(endpoint, "https"),
	})
	if (err != nil) {
		return nil, err
	}
	objectName := strings.TrimPrefix(strings.TrimRight(prefix, "/")+"/"+uploadedName, "/")
	if _, err := client.FPutObject(context.Background(), bucket, objectName, sourcePath, minio.PutObjectOptions{}); err != nil {
		return nil, err
	}
	dumpUrl := strings.TrimRight(endpoint, "/") + "/" + bucket + "/" + objectName
	return buildFileRecord(s.dumpMetadata, sourcePath, uploadedName, dumpUrl)
}

//通过SCP协议转储到另一主机
type ScpDumper struct {
	dumpMetadata *backupEntity.BackupDump
}

// 转储文件
func (s *ScpDumper) Dump(sourcePath, uploadedName string) (*backupEntity.BackupFileRecord, error) {
	if s.dumpMetadata.Credential == nil {
		return nil, errors.New("scp转储认证信息缺失")
	}
	credential, ok := (*s.dumpMetadata.Credential).(*backupEntity.SCPDumpCredential)
	if !ok {
		return nil, errors.New("scp转储认证信息无效")
	}
	if credential.User == "root" {
		return nil, errors.New("scp转储不允许使用root用户")
	}
	if credential.User == "" {
		return nil, errors.New("scp转储用户名不能为空")
	}
	host, port, remotePath, err := parseScpTarget(s.dumpMetadata.UrlTemplate)
	if (err != nil) {
		return nil, err
	}
	client, err := dialSSH(host, port, credential)
	if (err != nil) {
		return nil, err
	}
	defer client.Close()
	//上传到 urlTemplate 指定的远端目录
	if err := scpSend(client, sourcePath, uploadedName, remotePath); err != nil {
		return nil, err
	}
	remoteFile := uploadedName
	if remotePath != "." && remotePath != "" {
		remoteFile = strings.TrimRight(remotePath, "/") + "/" + uploadedName
	}
	dumpUrl := fmt.Sprintf("scp://%s@%s/%s", credential.User, net.JoinHostPort(host, port), remoteFile)
	return buildFileRecord(s.dumpMetadata, sourcePath, uploadedName, dumpUrl)
}

// 转储到本机另一目录
type LocalDumper struct {
	dumpMetadata *backupEntity.BackupDump
}

// 转储文件
func (s *LocalDumper) Dump(sourcePath, uploadedName string) (*backupEntity.BackupFileRecord, error) {
	if s.dumpMetadata.UrlTemplate == "" {
		return nil, errors.New("local转储路径前缀为空")
	}
	if err := os.MkdirAll(s.dumpMetadata.UrlTemplate, 0755); err != nil {
		return nil, err
	}
	destPath := filepath.Join(s.dumpMetadata.UrlTemplate, uploadedName)
	if err := copyFile(sourcePath, destPath); err != nil {
		return nil, err
	}
	return buildFileRecord(s.dumpMetadata, sourcePath, uploadedName, destPath)
}

// buildFileRecord 根据转储结果构造备份文件记录
func buildFileRecord(dumpMetadata *backupEntity.BackupDump, sourcePath, uploadedName, dumpUrl string) (*backupEntity.BackupFileRecord, error) {
	info, err := os.Stat(sourcePath)
	if (err != nil) {
		return nil, err
	}
	now := time.Now()
	record := &backupEntity.BackupFileRecord{
		MaidId:         dumpMetadata.MaidId,
		DumpId:         dumpMetadata.DumpId,
		OriginalPath:   sourcePath,
		FileName:       uploadedName,
		FileSize:       uint64(info.Size()),
		FileType:       "file",
		FileCreateTime: info.ModTime(),
		FileModifyTime: info.ModTime(),
		DumpUrl:        dumpUrl,
		CreateTime:     now,
		UpdateTime:     now,
	}
	// 若配置了保存时长,计算过期时间
	if dumpMetadata.HoursToLive != nil && *dumpMetadata.HoursToLive > 0 {
		record.ExpireTime = now.Add(time.Duration(*dumpMetadata.HoursToLive) * time.Hour)
	}
	return record, nil
}

// copyFile 将 src 文件复制到 dst
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if (err != nil) {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if (err != nil) {
		return err
	}
	defer out.Close()
	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Sync()
}

// parseScpTarget 解析 scp://host[:port][/remotePath] 形式的存储地址
func parseScpTarget(urlTemplate string) (host, port, remotePath string, err error) {
	u, err := url.Parse(urlTemplate)
	if (err != nil) {
		return "", "", "", err
	}
	if u.Scheme != "scp" {
		return "", "", "", fmt.Errorf("无效的scp存储地址: %s", urlTemplate)
	}
	host = u.Hostname()
	port = u.Port()
	if port == "" {
		port = "22"
	}
	remotePath = u.Path
	if remotePath == "" || remotePath == "/" {
		// 未指定路径时默认上传到远端 home 目录
		remotePath = "."
	} else {
		remotePath = strings.TrimPrefix(remotePath, "/")
	}
	return host, port, remotePath, nil
}

// dialSSH 使用密钥认证建立ssh连接
func dialSSH(host, port string, credential *backupEntity.SCPDumpCredential) (*ssh.Client, error) {
	signer, err := loadSigner(credential.Secret)
	if (err != nil) {
		return nil, err
	}
	config := &ssh.ClientConfig{
		User:            credential.User,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // TODO: 生产环境应校验主机密钥
		Timeout:         30 * time.Second,
	}
	return ssh.Dial("tcp", net.JoinHostPort(host, port), config)
}

// loadSigner 加载私钥,支持私钥原文或私钥文件路径
func loadSigner(secret string) (ssh.Signer, error) {
	// 若 secret 指向存在的文件,按文件路径读取
	if _, err := os.Stat(secret); err == nil {
		keyBytes, readErr := os.ReadFile(secret)
		if readErr != nil {
			return nil, readErr
		}
		return ssh.ParsePrivateKey(keyBytes)
	}
	// 否则按私钥原文解析
	return ssh.ParsePrivateKey([]byte(secret))
}

// scpSend 通过scp协议将本地文件上传到远端 remotePath 目录
func scpSend(client *ssh.Client, sourcePath, uploadedName, remotePath string) error {
	session, err := client.NewSession()
	if (err != nil) {
		return err
	}
	defer session.Close()
	stdin, err := session.StdinPipe()
	if (err != nil) {
		return err
	}
	stdout, err := session.StdoutPipe()
	if (err != nil) {
		return err
	}
	// scp -t <remotePath> 表示接收文件到远端指定目录("." 表示 home 目录)
	if err := session.Start("scp -t " + remotePath); err != nil {
		return err
	}

	file, err := os.Open(sourcePath)
	if (err != nil) {
		return err
	}
	defer file.Close()
	info, err := file.Stat()
	if (err != nil) {
		return err
	}

	// 读取远端首条确认
	if err := readScpAck(stdout); err != nil {
		return err
	}
	// 发送文件头: C<mode> <size> <filename>
	if _, err := fmt.Fprintf(stdin, "C%04o %d %s\n", info.Mode().Perm(), info.Size(), uploadedName); err != nil {
		return err
	}
	if err := readScpAck(stdout); err != nil {
		return err
	}
	// 发送文件内容
	if _, err := io.Copy(stdin, file); err != nil {
		return err
	}
	if _, err := stdin.Write([]byte{0}); err != nil {
		return err
	}
	if err := readScpAck(stdout); err != nil {
		return err
	}
	// 结束传输
	if _, err := stdin.Write([]byte("E\n")); err != nil {
		return err
	}
	if err := stdin.Close(); err != nil {
		return err
	}
	return session.Wait()
}

// readScpAck 读取scp协议中的单字节确认,0表示成功
func readScpAck(stdout io.Reader) error {
	buf := make([]byte, 1)
	if _, err := io.ReadFull(stdout, buf); err != nil {
		return err
	}
	if buf[0] == 0 {
		return nil
	}
	// 出错时远端会附带一行错误信息
	msg := make([]byte, 0, 256)
	tmp := make([]byte, 1)
	for {
		if _, err := io.ReadFull(stdout, tmp); err != nil {
			break
		}
		if tmp[0] == '\n' {
			break
		}
		msg = append(msg, tmp[0])
	}
	return fmt.Errorf("scp传输错误: %s", string(msg))
}

// parseS3Target 解析 s3 存储地址,返回 endpoint、bucket 和 key 前缀
// 支持两种风格:
//   - 虚拟主机风格(如 https://yes-my-admin.s3.us-east-1.amazonaws.com/backup):bucket 嵌在主机名中
//   - 路径风格(如 https://host[:port]/bucket/prefix):首段路径为 bucket
func parseS3Target(urlTemplate string) (endpoint, bucket, prefix string, err error) {
	u, err := url.Parse(urlTemplate)
	if (err != nil) {
		return "", "", "", err
	}
	if u.Scheme != "https" && u.Scheme != "http" {
		return "", "", "", fmt.Errorf("无效的s3存储地址: %s", urlTemplate)
	}
	trimmed := strings.Trim(u.Path, "/")
	var segments []string
	if trimmed != "" {
		segments = strings.Split(trimmed, "/")
	}

	// 虚拟主机风格: <bucket>.s3.<region>...,bucket 嵌在主机名中,需剥离
	if idx := strings.Index(u.Hostname(), ".s3."); idx > 0 {
		bucket = u.Hostname()[:idx]
		restHost := u.Hostname()[idx+1:] // s3.<region>...
		endpoint = u.Scheme + "://" + restHost
		prefix = strings.Join(segments, "/")
		return endpoint, bucket, prefix, nil
	}

	// 路径风格: 首段路径为 bucket,其余为 key 前缀
	if len(segments) == 0 {
		return "", "", "", fmt.Errorf("s3存储地址缺少bucket: %s", urlTemplate)
	}
	bucket = segments[0]
	prefix = strings.Join(segments[1:], "/")
	endpoint = u.Scheme + "://" + u.Host
	return endpoint, bucket, prefix, nil
}
