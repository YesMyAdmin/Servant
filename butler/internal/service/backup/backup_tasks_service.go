package backup

import (
	backupdto "butler/internal/model/dto/backup"
	"common/public/pkg"
	"butler/internal/repository/backup"
	"butler/internal/model/entity/backup"
	"math"
)

// 新建备份定时任务
func NewBackupTask(req *backupdto.NewBackupTaskReq) error {
	err := backup.NewBackupTask(entity.NewReqToPO(req))
	if err != nil {
		return err
	}
	return nil
}

// 更新备份定时任务
func EditBackupTask(req *backupdto.EditBackupTaskReq) error {
	err := backup.EditBackupTask(entity.EditReqToPO(req))
	if err != nil {
		return err
	}
	return nil
}

// 查询定时任务
func ListTasks(req *backupdto.ListTasksReq) (*pkg.PageableResp[backupdto.ListTasksResp], error) {
	tasks, total, err := backup.ListTasks(req.PageNum, req.PageSize, req.TaskName)
	if err != nil {
		return nil, pkg.InternalServerError(err.Error())
	}

	// 模型转 DTO
	contents := make([]backupdto.ListTasksResp, 0, len(tasks))
	for _, t := range tasks {
		resp := entity.ToListTasksResp(&t)
		contents = append(contents, *resp)
	}

	pages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &pkg.PageableResp[backupdto.ListTasksResp]{
		Total:    total,
		Pages:    pages,
		Contents: contents,
	}, nil
}

// 任务开关切换
func BackupTaskSwitch(taskId uint64, enabled bool) error {
	return backup.SwitchBackupTask(taskId, enabled)
}

// 删除任务
func DeleteBackupTask(taskId uint64) error {
	return backup.DeleteBackupTask(taskId)
}