package backup

import "time"
import "github.com/go-co-op/gocron"
// import 	backupdto "common/public/model/dto/backup"
// import backupsvc "common/public/service/backup"

// RegisterTask 注册任务
func RegisterTask(cron string) {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	scheduler := gocron.NewScheduler(timezone)
	scheduler.Cron(cron).Do(func() {
	})
}

// // DoBackup 执行备份任务
// func DoBackup(backupTask *backupdto.EditBackupTaskReq) {
// 	backupsvc.

// }