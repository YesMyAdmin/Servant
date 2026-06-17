package model

import "time"

// BackupTaskPO 备份任务数据库模型
type BackupTaskPO struct {
	FileId         int64     `gorm:"column:file_id;primaryKey"`
	MaidId         int64     `gorm:"column:maid_id"`
	MaidName       string    `gorm:"column:maid_name"`
	OriginalPath   string    `gorm:"column:original_path"`
	FileSize       uint64    `gorm:"column:file_size"`
	FileModifyTime time.Time `gorm:"column:file_modify_time"`
	VersionHash    uint64    `gorm:"column:version_hash"`
	RecordTime     time.Time `gorm:"column:record_time"`
	UpdateTime     time.Time `gorm:"column:update_time"`
}

// TableName 指定表名
func (BackupTaskPO) TableName() string {
	return "backup_tasks"
}
