package backup

import "time"

// BackupDumpPO 备份文件存储方式表 PO (对应表 backup_dumps)
type BackupDumpPO struct {
	DumpId      uint64     `gorm:"column:dump_id;primaryKey;comment:存储位置id" json:"dumpId"`
	DumpName    string     `gorm:"column:dump_name;type:varchar(32);not null;unique;comment:存储方式名称" json:"dumpName"`
	MaidId      uint64     `gorm:"column:maid_id;not null;comment:女仆节点id" json:"maidId"`
	DumpType    string     `gorm:"column:dump_type;type:varchar(16);not null;comment:存储类型(local:本地文件系统 scp:异地主机(scp协议连接) s3:对象存储s3协议)" json:"dumpType"`
	UrlTemplate string     `gorm:"column:url_template;type:varchar(1024);not null;comment:存储位置url模板" json:"urlTemplate"`
	HoursToLive int32      `gorm:"column:hours_to_live;not null;comment:保存时间(以小时为单位),过期由定时任务清理文件" json:"hoursToLive"`
	Credential  *string    `gorm:"column:credential;type:json;comment:认证信息,不同的存储方式有不同的认证信息" json:"credential"`
	Enabled     bool       `gorm:"column:enabled;type:tinyint(1);not null;default:1;comment:是否启用(1:是 0:否)" json:"enabled"`
	DeletedTime *time.Time `gorm:"column:deleted_time;default:null;comment:删除时间(留空为未删除,有时间为软删除)" json:"deletedTime"`
	CreateTime  time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:记录创建时间" json:"createTime"`
	OwnerId     uint64     `gorm:"column:owner_id;not null;comment:所有者id(默认创建者)" json:"ownerId"`
	UpdateTime  time.Time  `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:记录修改时间" json:"updateTime"`
}

// TableName 指定表名
func (BackupDumpPO) TableName() string {
	return "backup_dumps"
}
