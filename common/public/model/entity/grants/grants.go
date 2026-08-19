package grants

import (
	"time"
	"common/internal/model/po/grants"
)

//授权对象类型 (user/group/maid/logged_in/anonymous/global)
type AuthorizedTargetType string

const (
	User        AuthorizedTargetType = "user"       // 用户
	Group       AuthorizedTargetType = "group"      // 组
	Maid        AuthorizedTargetType = "maid"       // 女仆节点
	LoggedIn    AuthorizedTargetType = "logged_in" // 已登录用户
	Anonymous   AuthorizedTargetType = "anonymous"  // 匿名用户
	Global      AuthorizedTargetType = "global"     // 全局
)
//被授权访问的类型(page/api/file/datatable)
type AuthorizedContentType string

const (
	Page      	AuthorizedContentType = "page"      // 页面
	Api 		AuthorizedContentType = "api"	  // 接口
	File      	AuthorizedContentType = "file"      // 文件
	Datatable 	AuthorizedContentType = "datatable" // 数据表
)

// Grants 授权信息
type Grants struct {
	// GrantId 授权ID
	GrantId               uint64    
	// GrantName 授权名称
	GrantName             string     
	// AuthorizedContentType 被授权访问的类型
	AuthorizedContentType AuthorizedContentType     
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
		AuthorizedContentType: AuthorizedContentType(po.AuthorizedContentType),
		AuthorizedContent:     po.AuthorizedContent,
		Enabled:               po.Enabled,
		DeletedTime:           po.DeletedTime,
		CreateTime:            po.CreateTime,
		OwnerId:               po.OwnerId,
		UpdateTime:            po.UpdateTime,
	}
}