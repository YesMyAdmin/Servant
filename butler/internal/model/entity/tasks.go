package entity

import (
	"butler/internal/model/dto"
	"butler/internal/model/po"
	"time"
)

// 任务
type Tasks struct {
	// 任务ID
	TaskId uint64
	// 女仆节点id
	MaidId uint64
	// 定时任务名称
	TaskName string
	// 任务类型(builtin:系统内置的任务 backup:备份任务)
	TaskType string
	// 触发器类型(cron:cron表达式)
	Trigger string
	// CRON表达式
	Cron string
	// 是否启用
	Enabled bool
	// 删除时间
	DeletedTime time.Time
	//创建时间
	CreateTime time.Time
	// 所有者id
	OwnerId uint64
	// 记录修改时间
	UpdateTime time.Time
}

// BackupTask 备份任务
type BackupTask struct {
	Tasks
	// 备份模式(full:全量 incremental:增量)
	Mode string
	// 需要备份的文件/文件夹路径
	Source string
}

func Load(po *po.BackupTaskPO) *BackupTask {
	return &BackupTask{
		Tasks: Tasks{
			TaskId: po.TaskId,
		},
		Mode: po.Mode,
		Source: po.Source,
	}
}

func DumpPO(entity *BackupTask) {

}

func DumpListTasksResp(entity *BackupTask) *dto.ListTasksResp {
	return &dto.ListTasksResp{
		TaskId: entity.TaskId,
		MaidId: entity.MaidId,
		Mode: entity.Mode,
		Source: entity.Source,
		Cron: entity.Cron,
		Enabled: entity.Enabled,
		OwnerId: entity.OwnerId,
		UpdateTime: entity.UpdateTime,
	}
}
