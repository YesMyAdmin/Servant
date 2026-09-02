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
	//查询文件
	if (nil == backupTask) {
		return nil
	}
	//备份记录id
	backupId := 0
	if (backupTask.Mode == backupEntity.Full) {
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
			temporaryZipFile := "./cached/backup/" + file.Name() + "_" + strconv.Itoa(backupId) + ".zip"
			err := utils.ZipFolder(backupTask.Source, temporaryZipFile)
			if (err != nil) {
				return err
			}
		} else {
			//如果目标是文件,直接上传
		}
	} else {
		//增量备份模式
	}
	return nil
}

// 备份文件转储抽象接口
type BackupDumper interface {
	// Upload 上传文件
	Dump(sourcePath, uploadedName string)
}

// S3协议转储器
type S3Dumper struct {
	DumpUrl string
}

// 转储文件
func (s *S3Dumper) Dump(sourcePath, uploadedName string) {

}

//通过SCP协议转储到另一主机
type SCPDumper struct {

}

// 转储文件
func (s *SCPDumper) Dump(sourcePath, uploadedName string) {
}

// 转储到本机另一目录
type LocalDumper struct {

}

// 转储文件
func  (s *LocalDumper) Dump(sourcePath, uploadedName string) {

}