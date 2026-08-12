package grants

import (
	"time"
	"common/internal/model/po/grants"
)

// Grants 授权信息
type Grants struct {
	// GrantId 授权ID
	GrantId               uint64    
	// GrantName 授权名称
	GrantName             string     
	// AuthorizedTargetType 授权对象类型
	AuthorizedTargetType  string     
	// AuthorizedTargetId 被授权对象id
	AuthorizedTargetId    *uint64    
	// AuthorizedContentType 被授权访问的类型
	AuthorizedContentType string     
	// AuthorizedContent 被授权内容
	AuthorizedContent     string     
	// Enabled 是否启用
	Enabled               string
	// DeletedTime 删除时间     
	DeletedTime           *time.Time
	// CreateTime 创建时间
	CreateTime            time.Time
	// OwnerId 所有者id  
	OwnerId               int64
	// UpdateTime 修改时间      
	UpdateTime            time.Time  
}

// LoadFromPO 将PO转换为实体
func LoadFromPO(po *grants.GrantPO) *Grants {
	if po == nil {
		return nil
	}
	return &Grants{
		GrantId:               po.GrantId,
		GrantName:             po.GrantName,
		AuthorizedTargetType:  po.AuthorizedTargetType,
		AuthorizedTargetId:    po.AuthorizedTargetId,
		AuthorizedContentType: po.AuthorizedContentType,
		AuthorizedContent:     po.AuthorizedContent,
		Enabled:               po.Enabled,
		DeletedTime:           po.DeletedTime,
		CreateTime:            po.CreateTime,
		OwnerId:               po.OwnerId,
		UpdateTime:            po.UpdateTime,
	}
}