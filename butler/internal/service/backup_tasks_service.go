package service

import (
	"butler/internal/dto"
	"butler/internal/model"
	"butler/internal/pkg"
	"butler/internal/repository"
	"math"
)

//新建备份定时任务
func NewBackupTask(req *dto.NewBackupTaskReq) error {
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

// modelToDto 将数据库模型转换为响应 DTO
func modelToDto(t model.BackupTaskPO) dto.ListTasksResp {
	return dto.ListTasksResp{
		FileId:         t.FileId,
		MaidId:         t.MaidId,
		MaidName:       t.MaidName,
		OriginalPath:   t.OriginalPath,
		FileSize:       t.FileSize,
		FileModifyTime: t.FileModifyTime,
		VersionHash:    t.VersionHash,
		RecordTime:     t.RecordTime,
		UpdateTime:     t.UpdateTime,
	}
}
