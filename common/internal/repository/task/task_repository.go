package task

import (
	taskPo "common/internal/model/po/task"
	"common/public/database"
)

// SelectById 根据任务ID查询定时任务
func SelectById(taskId uint64) (*taskPo.TaskPO, error) {
	db := database.DB.Model(&taskPo.TaskPO{}).
		Where("task_id = ?", taskId).
		Where("deleted_time IS NULL")
	var task taskPo.TaskPO
	err := db.First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// ListTasks 根据可选条件查询定时任务
// 各条件均为可选, 传入零值表示不按该条件过滤:
//   - maidId: 女仆节点id, 传 0 表示不过滤
//   - taskName: 任务名称(模糊匹配), 传空字符串表示不过滤
//   - taskType: 任务类型, 传空字符串表示不过滤
func ListTasks(maidId uint64, taskName, taskType string) ([]taskPo.TaskPO, error) {
	db := database.DB.Model(&taskPo.TaskPO{}).Where("deleted_time IS NULL")

	if maidId != 0 {
		db = db.Where("maid_id = ?", maidId)
	}
	if taskName != "" {
		db = db.Where("instr(task_name, ?)", taskName)
	}
	if taskType != "" {
		db = db.Where("task_type = ?", taskType)
	}

	var tasks []taskPo.TaskPO
	if err := db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
