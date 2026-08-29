package backup

import (
	"common/internal/model/po/task"
)

// BackupTaskPO 备份任务表 PO (对应表 backup_tasks)
type BackupTaskPO struct {
	Mode        string     `gorm:"column:mode;type:varchar(12);not null;comment:备份模式(full:全量 incremental:增量)" json:"mode"`
	Source      string     `gorm:"column:source;type:varchar(256);not null;comment:需要备份的文件/文件夹路径" json:"source"`
	task.TaskPO
}

// TableName 指定表名
func (BackupTaskPO) TableName() string {
	return "backup_tasks"
}
