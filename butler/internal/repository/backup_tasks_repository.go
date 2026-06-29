package repository

import (
	"butler/internal/database"
)

// ListTasks 分页查询备份任务，支持按任务名称模糊搜索
func ListTasks(pageNum, pageSize int, taskName string) ([]model.BackupTaskPO, int64, error) {
	db := database.DB.Model(&model.BackupTaskPO{})

	// 按名称模糊搜索
	if taskName != "" {
		db = db.Where("instr(task_name, ?)", taskName)
	}

	// 先查询总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	var tasks []model.BackupTaskPO
	offset := (pageNum - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, total, nil
}

// 添加新的备份任务
func NewBackupTask(backupTask *model.BackupTaskPO) error {
	db := database.DB.Model(&model.BackupTaskPO{})
	err := db.Create(backupTask).Error
	if (err != nil) {
		return err
	}
	return nil
}