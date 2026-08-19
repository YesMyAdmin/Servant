package group

import (
	groupPo "common/internal/model/po/group"
	"common/public/database"
	"time"
)

// SelectList 分页查询用户组
func SelectList(pageNum, pageSize int) (*[]groupPo.GroupPO, int64, error) {
	db := database.DB.Model(&groupPo.GroupPO{}).Where("deleted_time IS NULL")

	// 先查询总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	var groups []groupPo.GroupPO
	offset := (pageNum - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Find(&groups).Error; err != nil {
		return nil, 0, err
	}

	return &groups, total, nil
}

// SelectById 根据用户组ID查询用户组信息
func SelectById(groupId uint64) (*groupPo.GroupPO, error) {
	db := database.DB.Model(&groupPo.GroupPO{}).
		Where("group_id = ?", groupId).
		Where("deleted_time IS NULL")
	var group groupPo.GroupPO
	err := db.First(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

// Update 修改用户组信息
func Update(group *groupPo.GroupPO) error {
	db := database.DB.Model(&groupPo.GroupPO{})
	err := db.Where("group_id = ?", group.GroupId).Updates(group).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete 逻辑删除用户组(软删除)
func Delete(groupId uint64) error {
	db := database.DB.Model(&groupPo.GroupPO{})
	err := db.Where("group_id = ?", groupId).UpdateColumn("deleted_time", time.Now()).Error
	if err != nil {
		return err
	}
	return nil
}
