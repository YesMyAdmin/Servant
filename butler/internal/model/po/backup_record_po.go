package po

import "time"

// BackupRecordPO 备份文件记录表 PO (对应表 backup_records)
type BackupRecordPO struct {
	BackupRecordId       uint64     `gorm:"column:backup_record_id;primaryKey;comment:记录id" json:"backupRecordId"`
	FileId               uint64     `gorm:"column:file_id;not null;comment:文件id，使用maidId+originalPath计算得出，用于追踪备份历史" json:"fileId"`
	FileName             string     `gorm:"column:file_name;type:varchar(128);not null;comment:文件名称" json:"fileName"`
	MaidId               uint64     `gorm:"column:maid_id;not null;comment:女仆节点id" json:"maidId"`
	TaskId               uint64     `gorm:"column:task_id;not null;comment:任务id" json:"taskId"`
	OriginalPath         string     `gorm:"column:original_path;type:varchar(1024);not null;comment:源文件路径" json:"originalPath"`
	FileSize             uint64     `gorm:"column:file_size;not null;comment:文件大小(bytes)" json:"fileSize"`
	FileCreateTime       time.Time  `gorm:"column:file_create_time;not null;comment:文件创建时间" json:"fileCreateTime"`
	FileModifyTime       time.Time  `gorm:"column:file_modify_time;not null;comment:文件修改时间" json:"fileModifyTime"`
	VersionHash          uint64     `gorm:"column:version_hash;not null;comment:文件版本hash,由xxhash64生成" json:"versionHash"`
	DumpId               uint64     `gorm:"column:dump_id;not null;comment:转储位置id" json:"dumpId"`
	ExpireTime           time.Time  `gorm:"column:expire_time;not null;comment:过期时间" json:"expireTime"`
	FileActualDeletedTime *time.Time `gorm:"column:file_actual_deleted_time;default:null;comment:文件删除时间(留空为未删除,有时间为已删除)" json:"fileActualDeletedTime"`
	DeletedTime          *time.Time `gorm:"column:deleted_time;default:null;comment:删除时间(留空为未删除,有时间为软删除)" json:"deletedTime"`
	CreateTime           time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:记录创建时间" json:"createTime"`
	OwnerId              uint64     `gorm:"column:owner_id;not null;comment:所有者id(默认创建者)" json:"ownerId"`
	UpdateTime           time.Time  `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:记录修改时间" json:"updateTime"`
}

// TableName 指定表名
func (BackupRecordPO) TableName() string {
	return "backup_records"
}