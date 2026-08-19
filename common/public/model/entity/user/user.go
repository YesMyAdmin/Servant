package user

import (
	"common/public/model/entity/grants"
	"common/public/model/entity/group"
	userPO "common/internal/model/po/user"
	grantsPO "common/internal/model/po/grants"
	groupPO "common/internal/model/po/group"
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
	Description string
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

// LoadFromPO 从用户信息PO加载
func LoadFromPO(userPO *userPO.UserPO, groupPO *[]groupPO.GroupPO, grantsPO *[]grantsPO.GrantPO ) *User {
	user := &User{}
	user.UserId = userPO.UserId
	user.Name = userPO.Name
	user.Password = userPO.Password
	user.Email = userPO.Email
	user.Phone = userPO.Phone
	user.FreezeTime = userPO.FreezeTime
	user.LastActiveTime = userPO.LastActiveTime
	user.InviterUserId = userPO.InviterUserId
	user.Description = userPO.Description
	user.DeletedTime = userPO.DeletedTime
	user.CreateTime = userPO.CreateTime
	user.OwnerId = userPO.OwnerId
	user.UpdateTime = userPO.UpdateTime
	user.Groups = &[]group.Group{}
	for _, po := range *groupPO {
		group := group.LoadFromPO(&po)
		*user.Groups = append(*user.Groups, *group)
	}
	user.UserGrants = &[]grants.Grants{}
	for _, po := range *grantsPO {
		grant := grants.LoadFromPO(&po)
		*user.UserGrants = append(*user.UserGrants, *grant)
	}
	return user
}