package user

import "time"

// UserPO 用户表 PO (对应表 users)
type UserPO struct {
	UserId         uint64     `gorm:"column:user_id;primaryKey;comment:用户id" json:"userId"`
	Name           string     `gorm:"column:name;type:varchar(64);not null;comment:用户名" json:"name"`
	Password       string     `gorm:"column:password;type:varchar(128);not null;comment:密码(已混淆)" json:"password"`
	Salt           int        `gorm:"column:salt;type:int;not null;comment:盐" json:"salt"`
	OtpProvider    string     `gorm:"column:otp_provider;type:varchar(16);not null;default:none;comment:otp方法" json:"otpProvider"`
	OtpSecret      string     `gorm:"column:otp_secret;type:varchar(128);comment:otp密钥(未加密)" json:"otpSecret"`
	Email          string     `gorm:"column:email;type:varchar(128);comment:邮件(未加密)" json:"email"`
	Phone          string     `gorm:"column:phone;type:varchar(32);comment:手机号(未加密)" json:"phone"`
	FreezeTime     time.Time  `gorm:"column:freeze_time;default:null;comment:冻结时间" json:"freezeTime"`
	LastActiveTime time.Time  `gorm:"column:last_active_time;default:null;comment:用户上次活动时间" json:"lastActiveTime"`
	InviterUserId  uint64     `gorm:"column:inviter_user_id;default:0;comment:用户的邀请者id" json:"inviterUserId"`
	Description    string     `gorm:"column:description;type:varchar(255);comment:描述" json:"description"`
	DeletedTime    *time.Time `gorm:"column:deleted_time;default:null;comment:删除时间(留空为未删除,有时间为软删除)" json:"deletedTime"`
	CreateTime     time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:记录创建时间" json:"createTime"`
	OwnerId        int64      `gorm:"column:owner_id;not null;comment:所有者id(默认创建者)" json:"ownerId"`
	UpdateTime     time.Time  `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:记录修改时间" json:"updateTime"`
}

// TableName 指定表名
func (UserPO) TableName() string {
	return "users"
}
