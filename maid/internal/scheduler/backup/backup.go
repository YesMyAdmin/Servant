package backup

import (
	backupEntity "common/public/model/entity/backup"
	backupsvc "common/public/service/backup"
	"common/public/utils"
	"os"
	"strconv"
	"time"
	"github.com/go-co-op/gocron"
)

// RegisterTask 注册任务
func RegisterTask(cron string) {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	scheduler := gocron.NewScheduler(timezone)
	scheduler.Cron(cron).Do(func() {
	})
}

// DoBackup 执行备份任务
func DoBackup(taskId uint64) error {
	//加载备份任务信息
	backupTask, err:= backupsvc.GetBackupTask(taskId)
	if (err != nil) {
		return err
	}
	//获取备份存储记录,构造备份存储器列表
	//一个备份任务会下挂多个备份存储，需要将同一文件依次上传
	dumps, listDumpsErr := backupsvc.ListDumpsByTaskId(taskId)
	if (listDumpsErr != nil) {
		return listDumpsErr
	}
	dumpers := make([]BackupDumper, len(dumps))
	for _,dump := range dumps {
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
	if (fileInfo.IsDir()) {
		//如果目标是文件夹,将文件压缩后上传
		fileName := file.Name() + "_" + strconv.Itoa(backupId) + ".zip"
		temporaryZipFile := "./cached/backup/" + fileName
		err := utils.ZipFolder(backupTask.Source, temporaryZipFile)
		if (err != nil) {
			return err
		}	
		for _, dumper := range dumpers {
			dumper.Dump(temporaryZipFile, fileName)
		}
	} else {
		//如果目标是文件,直接上传
		fileName := file.Name() + "_" + strconv.Itoa(backupId) + ".zip"
		for _, dumper := range dumpers {
			dumper.Dump(backupTask.Source, fileName)
		}
	}
	return nil
}

// 获取一个转储器实例
func makeDumper(dumpMetadata *backupEntity.BackupDump) (BackupDumper) {
	switch dumpMetadata.DumpType {
		case backupEntity.S3:
			return &S3Dumper{dumpMetadata: dumpMetadata}
		case backupEntity.Scp:
			return &ScpDumper{dumpMetadata: dumpMetadata}
		default:
			return &LocalDumper{dumpMetadata: dumpMetadata}
	}
}

// 备份文件转储抽象接口
type BackupDumper interface {
	// Upload 上传文件
	Dump(sourcePath, uploadedName string) (*backupEntity.BackupFileRecord)
}

// S3协议转储器
type S3Dumper struct {
	dumpMetadata *backupEntity.BackupDump
}

// 转储文件
func (s *S3Dumper) Dump(sourcePath, uploadedName string) (*backupEntity.BackupFileRecord) {
	return nil
}

//通过SCP协议转储到另一主机
type ScpDumper struct {
	dumpMetadata *backupEntity.BackupDump
}

// 转储文件
func (s *ScpDumper) Dump(sourcePath, uploadedName string) (*backupEntity.BackupFileRecord) {
	return nil
}

// 转储到本机另一目录
type LocalDumper struct {
	dumpMetadata *backupEntity.BackupDump
}

// 转储文件
func  (s *LocalDumper) Dump(sourcePath, uploadedName string) (*backupEntity.BackupFileRecord) {
	return nil
}