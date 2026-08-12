package grants

import (
	"common/public/database"
	grantsPo "common/internal/model/po/grants"
	grantsConst "common/public/consts/grants"
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
func SelectByContentType(authorizedContentType grantsConst.AuthorizedContentType) (*[]grantsPo.GrantPO, error) {
	db := database.DB.Model(&grantsPo.GrantPO{}).Where("authorized_content_type = ?", authorizedContentType)
	var grants []grantsPo.GrantPO
	err := db.Find(&grants).Error
	if err != nil {
		return nil, err
	}
	return &grants, nil
}

