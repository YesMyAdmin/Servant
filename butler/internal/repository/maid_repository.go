package repository

import (
	"butler/internal/database"
	"butler/internal/model/po"
	"time"
)

// ListMaids 分页查询女仆节点，支持按名称模糊搜索
func ListMaids(pageNum, pageSize int, maidName string) ([]po.MaidPO, int64, error) {
	db := database.DB.Model(&po.MaidPO{})

	// 按名称模糊搜索
	if maidName != "" {
		db = db.Where("instr(maid_name, ?)", maidName)
	}

	// 先查询总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	var maids []po.MaidPO
	offset := (pageNum - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Find(&maids).Error; err != nil {
		return nil, 0, err
	}

	return maids, total, nil
}

// NewMaid 添加新的女仆节点
func NewMaid(maid *po.MaidPO) error {
	db := database.DB.Model(&po.MaidPO{})
	err := db.Create(maid).Error
	if err != nil {
		return err
	}
	return nil
}

// EditMaid 更新女仆节点
func EditMaid(maid *po.MaidPO) error {
	db := database.DB.Model(&po.MaidPO{})
	err := db.Where("maid_id = ?", maid.MaidId).Updates(maid).Error
	if err != nil {
		return err
	}
	return nil
}

// SwitchMaid 切换女仆节点启用状态
func SwitchMaid(maidId uint64, enabled bool) error {
	db := database.DB.Model(&po.MaidPO{})
	err := db.Where("maid_id = ?", maidId).UpdateColumn("enabled", enabled).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteMaid 删除女仆节点(软删除)
func DeleteMaid(maidId uint64) error {
	db := database.DB.Model(&po.MaidPO{})
	err := db.Where("maid_id = ?", maidId).UpdateColumn("deleted_time", time.Now()).Error
	if err != nil {
		return err
	}
	return nil
}