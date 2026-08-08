package po

import "time"

// MaidPO 女仆节点表 PO (对应表 maids)
type MaidPO struct {
	MaidId      uint64     `gorm:"column:maid_id;primaryKey;comment:备份任务id"                                   json:"maidId"`
	MaidName    string     `gorm:"column:maid_name;type:varchar(32);not null;comment:女仆节点名称"                  json:"maidName"`
	HostPort    string     `gorm:"column:host_port;type:varchar(50);not null;comment:女仆节点的主机地址+端口"           json:"hostPort"`
	AccessSecret string     `gorm:"column:access_token;type:varchar(128);not null;comment:访问令牌"                  json:"accessSecret"`
	Enabled     int8       `gorm:"column:enabled;type:tinyint;not null;default:1;comment:是否启用(1:是 0:否)"         json:"enabled"`
	DeletedTime *time.Time `gorm:"column:deleted_time;default:null;comment:删除时间(留空为未删除,有时间为软删除)"           json:"deletedTime"`
	CreateTime  time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:记录创建时间"        json:"createTime"`
	OwnerId     int64      `gorm:"column:owner_id;not null;comment:所有者id(默认创建者)"                              json:"ownerId"`
	UpdateTime  time.Time  `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:记录修改时间" json:"updateTime"`
}

// TableName 指定表名
func (MaidPO) TableName() string {
	return "maids"
}