package group

import (
	"common/public/model/entity/grants"
	"time"
)

//Group 用户组
type Group struct {
	//用户组id
	GroupId uint64
	//用户组名称
	GroupName string
	DeletedTime           *time.Time
	// CreateTime 创建时间
	CreateTime            time.Time
	// OwnerId 所有者id  
	OwnerId               int64
	// UpdateTime 修改时间      
	UpdateTime            time.Time
	//授权信息
	Grants *[]grants.Grants
}