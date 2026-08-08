package po

import "time"

// BackupTaskPO 备份任务表 PO (对应表 backup_tasks)
type BackupTaskPO struct {
	TaskId      uint64      `gorm:"column:task_id;primaryKey;comment:任务id,与tasks表一致"      json:"taskId"`
	Mode        string     `gorm:"column:mode;type:varchar(12);not null;comment:备份模式(full:全量 incremental:增量)" json:"mode"`
	Source      string     `gorm:"column:source;type:varchar(256);not null;comment:需要备份的文件/文件夹路径" json:"source"`
	DeletedTime time.Time `gorm:"column:deleted_time;default:null;comment:删除时间(留空为未删除,有时间为软删除)" json:"deletedTime"`
	CreateTime  time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:记录创建时间" json:"createTime"`
	OwnerId     int64      `gorm:"column:owner_id;not null;comment:所有者id(默认创建者)" json:"ownerId"`
	UpdateTime  time.Time  `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:记录修改时间" json:"updateTime"`
}

// TableName 指定表名
func (BackupTaskPO) TableName() string {
	return "backup_tasks"
}
