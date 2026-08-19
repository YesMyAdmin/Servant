package grants

import "time"

// GrantPO 授权表 PO (对应表 grants)
type GrantPO struct {
	GrantId               uint64     `gorm:"column:grant_id;primaryKey;comment:记录id" json:"grantId"`
	GrantName             string     `gorm:"column:grant_name;type:varchar(32);not null;comment:权限名称" json:"grantName"`
	AuthorizedContentType string     `gorm:"column:authorized_content_type;type:varchar(10);not null;comment:被授权访问的类型(page/api/file/datatable)" json:"authorizedContentType"`
	AuthorizedContent     string     `gorm:"column:authorized_content;type:varchar(128);not null;comment:被授权内容,如GET:/panel/backup/tasks等" json:"authorizedContent"`
	Enabled               string     `gorm:"column:enabled;type:char(1);not null;default:1;comment:是否启用(1:是 0:否)" json:"enabled"`
	DeletedTime           *time.Time `gorm:"column:deleted_time;default:null;comment:删除时间(留空为未删除,有时间为软删除)" json:"deletedTime"`
	CreateTime            time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:记录创建时间" json:"createTime"`
	OwnerId               int64      `gorm:"column:owner_id;not null;comment:所有者id(默认创建者)" json:"ownerId"`
	UpdateTime            time.Time  `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:记录修改时间" json:"updateTime"`
}

// TableName 指定表名
func (GrantPO) TableName() string {
	return "grants"
}
