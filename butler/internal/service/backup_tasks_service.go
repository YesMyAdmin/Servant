package service

import (
	"butler/internal/model/dto"
	"common/public/pkg"
	"butler/internal/repository"
	"butler/internal/model/entity"
	"math"
)

// 新建备份定时任务
func NewBackupTask(req *dto.NewBackupTaskReq) error {
	err := repository.NewBackupTask(entity.NewReqToPO(req))
	if err != nil {
		return err
	}
	return nil
}

// 更新备份定时任务
func EditBackupTask(req *dto.EditBackupTaskReq) error {
	err := repository.EditBackupTask(entity.EditReqToPO(req))
	if err != nil {
		return err
	}
	return nil
}

// 查询定时任务
func ListTasks(req *dto.ListTasksReq) (*pkg.PageableResp[dto.ListTasksResp], error) {
	tasks, total, err := repository.ListTasks(req.PageNum, req.PageSize, req.TaskName)
	if err != nil {
		return nil, pkg.InternalServerError(err.Error())
	}

	// 模型转 DTO
	contents := make([]dto.ListTasksResp, 0, len(tasks))
	for _, t := range tasks {
		resp := entity.ToListTasksResp(&t)
		contents = append(contents, *resp)
	}

	pages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &pkg.PageableResp[dto.ListTasksResp]{
		Total:    total,
		Pages:    pages,
		Contents: contents,
	}, nil
}

// 任务开关切换
func BackupTaskSwitch(taskId uint64, enabled bool) error {
	return repository.SwitchBackupTask(taskId, enabled)
}

// 删除任务
func DeleteBackupTask(taskId uint64) error {
	return repository.DeleteBackupTask(taskId)
}


