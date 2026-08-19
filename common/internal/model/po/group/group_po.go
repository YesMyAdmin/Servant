package group

import "time"

// GroupPO 用户组表 PO (对应表 groups)
type GroupPO struct {
	GroupId     uint64     `gorm:"column:group_id;primaryKey;comment:组id" json:"groupId"`
	GroupName   string     `gorm:"column:group_name;type:varchar(32);not null;comment:组名" json:"groupName"`
	DeletedTime *time.Time `gorm:"column:deleted_time;default:null;comment:删除时间(留空为未删除,有时间为软删除)" json:"deletedTime"`
	CreateTime  time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:记录创建时间" json:"createTime"`
	OwnerId     int64      `gorm:"column:owner_id;not null;comment:所有者id(默认创建者)" json:"ownerId"`
	UpdateTime  time.Time  `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:记录修改时间" json:"updateTime"`
}

// TableName 指定表名
func (GroupPO) TableName() string {
	return "groups"
}
