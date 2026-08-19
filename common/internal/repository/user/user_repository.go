package user

import (
	userPo "common/internal/model/po/user"
	"common/public/database"
	"time"
)

// ListUsers 分页查询用户，支持按用户名模糊搜索
func ListUsers(pageNum, pageSize int, name string) (*[]userPo.UserPO, int64, error) {
	db := database.DB.Model(&userPo.UserPO{}).Where("deleted_time IS NULL")

	// 按用户名模糊搜索
	if name != "" {
		db = db.Where("instr(name, ?)", name)
	}

	// 先查询总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	var users []userPo.UserPO
	offset := (pageNum - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return &users, total, nil
}

// SelectById 根据用户ID查询用户信息
func SelectById(userId uint64) (*userPo.UserPO, error) {
	db := database.DB.Model(&userPo.UserPO{}).
		Where("user_id = ?", userId).
		Where("deleted_time IS NULL")
	var user userPo.UserPO
	err := db.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// EditUser 修改用户信息
func EditUser(user *userPo.UserPO) error {
	db := database.DB.Model(&userPo.UserPO{})
	err := db.Where("user_id = ?", user.UserId).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser 逻辑删除用户(软删除)
func DeleteUser(userId uint64) error {
	db := database.DB.Model(&userPo.UserPO{})
	err := db.Where("user_id = ?", userId).UpdateColumn("deleted_time", time.Now()).Error
	if err != nil {
		return err
	}
	return nil
}
