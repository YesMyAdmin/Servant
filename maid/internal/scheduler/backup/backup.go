package backup

import (
	backupsvc "common/public/service/backup"
	"log/slog"
	"time"

	"github.com/go-co-op/gocron"
)

//import backupEntity "common/public/model/entity/backup"

// RegisterTask 注册任务
func RegisterTask(cron string) {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	scheduler := gocron.NewScheduler(timezone)
	scheduler.Cron(cron).Do(func() {
	})
}

// DoBackup 执行备份任务
func DoBackup(taskId uint64) {
	//加载备份任务信息
	backupTask, err:= backupsvc.GetBackupTask(taskId)
	if (err != nil) {
		slog.Error("Load backup task %s failed", taskId, err)
	}
	//查询文件
	if (nil != backupTask) {
		
	}

}