package grants

import (
	grantsPo "common/internal/model/po/grants"
	"common/public/database"
)

// SelectAll 查询所有权限记录
func SelectAll() (*[]grantsPo.GrantPO, error) {
	db := database.DB.Model(&grantsPo.GrantPO{})
	var grants []grantsPo.GrantPO
	err := db.Find(&grants).Error
	if err != nil {
		return nil, err
	}
	return &grants, nil
}

// SelectByContentType 根据被授权访问的类型查询权限记录
// authorizedContentType: 被授权访问的类型，如page/api/file/datatable
func SelectByContentType(authorizedContentType string) (*[]grantsPo.GrantPO, error) {
	db := database.DB.Model(&grantsPo.GrantPO{}).Where("authorized_content_type = ?", authorizedContentType)
	var grants []grantsPo.GrantPO
	err := db.Find(&grants).Error
	if err != nil {
		return nil, err
	}
	return &grants, nil
}

// SelectByTargetId 根据授权对象类型和被授权对象id查找权限记录
// authorizedTargetType: 授权对象类型，如user/role/group
// authorizedTargetId: 授权对象id
func SelectByTargetId(authorizedTargetType string, authorizedTargetId uint64) (*[]grantsPo.GrantPO, error) {
	db := database.DB.Model(&grantsPo.GrantPO{}).Joins(
		"granted_targets", 
		database.DB.Select("grant_id").
		Where("authorized_target_type = ?", authorizedTargetType).
		Where("authorized_target_id = ?", authorizedTargetId),
	)
	var grants []grantsPo.GrantPO
	err := db.Find(&grants).Error
	if err != nil {
		return nil, err
	}
	return &grants, nil
}