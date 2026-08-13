package user

import (
	"common/public/model/entity/grants"
	"common/public/model/entity/group"
	"time"
)

type OtpProviderType string 
const (
	None OtpProviderType="none"
	Totp OtpProviderType="totp"
)

// User 用户
type User struct {
	//用户id
	UserId uint64
	//用户名
	Name string
	//密码(已混淆)
	Password string
	//盐
	Salt int
	//otp方法
	OtpProvider OtpProviderType
	//otp密钥(未加密)
	OtpSecret string
	//邮件(未加密)
	Email string
	//手机号(未加密)
	Phone string
	//冻结时间
	FreezeTime time.Time
	//用户上次活动时间
	LastActiveTime time.Time
	//用户的邀请者id
	InviterUserId uint64
	//描述
	Desc string
	// DeletedTime 删除时间     
	DeletedTime           *time.Time
	// CreateTime 创建时间
	CreateTime            time.Time
	// OwnerId 所有者id  
	OwnerId               int64
	// UpdateTime 修改时间      
	UpdateTime            time.Time  
	// Groups 用户所在的用户组
	Groups *[]group.Group
	// UserGrants 单独对此用户的授权信息
	UserGrants *[]grants.Grants 
}