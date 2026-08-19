package group

import (
	"common/internal/model/po/group"
	"common/public/model/entity/grants"
	"time"
)

//Group 用户组
type Group struct {
	//用户组id
	GroupId uint64
	//用户组名称
	GroupName string
	// DeletedTime 删除时间
	DeletedTime *time.Time
	// CreateTime 创建时间
	CreateTime time.Time
	// OwnerId 所有者id
	OwnerId int64
	// UpdateTime 修改时间
	UpdateTime time.Time
	//授权信息
	Grants *[]grants.Grants
}

// LoadFromPO 将PO转换为实体
func LoadFromPO(po *group.GroupPO) *Group {
	if po == nil {
		return nil
	}
	return &Group{
		GroupId:     po.GroupId,
		GroupName:   po.GroupName,
		DeletedTime: po.DeletedTime,
		CreateTime:  po.CreateTime,
		OwnerId:     po.OwnerId,
		UpdateTime:  po.UpdateTime,
	}
}
