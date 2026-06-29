package service

import (
	"butler/internal/dto"
	"butler/internal/pkg"
	"butler/internal/repository"
	"math"
	"time"
)





//新建备份定时任务
func NewBackupTask(req *dto.NewBackupTaskReq) error {

	repository.NewBackupTask(nil)
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
		contents = append(contents, modelToDto(t))
	}

	pages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &pkg.PageableResp[dto.ListTasksResp]{
		Total:    int(total),
		Pages:    pages,
		Contents: contents,
	}, nil
}
