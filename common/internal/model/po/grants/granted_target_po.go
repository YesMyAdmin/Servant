package grants

// GrantedTargetPO 被授权对象表 (对应表 granted_targets)
type GrantedTargetPO struct {
	//授权记录id
	GrantId         uint64     `gorm:"column:grant_id;not null;comment:授权记录id" json:"grantId"`
	//授权对象类型(user/group/maid/logged_in/anonymous/global)
	AuthorizedTargetType string   `gorm:"column:authorized_target_type;not null;comment:被授权对象类型" json:"authorizedTargetType"`
	// AuthorizedTargetId 被授权对象id(用户id或组id,匿名对象/全局设置留空)
	AuthorizedTargetId *uint64   `gorm:"column:authorized_target_id;not null;comment:被授权对象id" json:"authorizedTargetId"`
	Grants *GrantPO `gorm:"foreignKey:GrantId;references:GrantId" json:"grants"`
}