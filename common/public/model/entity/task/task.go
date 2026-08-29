package task

import (
	"time"

	taskPO "common/internal/model/po/task"
)

// TaskType 任务类型
type TaskType string

const (
	// Builtin 系统内置的任务
	Builtin TaskType = "builtin"
	// Backup 备份任务
	Backup TaskType = "backup"
)

// TriggerType 触发器类型
type TriggerType string

const (
	// Cron cron表达式触发器
	Cron TriggerType = "cron"
)

// Task 定时任务
type Task struct {
	// TaskId 任务ID
	TaskId uint64
	// MaidId 女仆节点id
	MaidId uint64
	// TaskName 定时任务名称
	TaskName string
	// TaskType 任务类型
	TaskType TaskType
	// Trigger 触发器类型
	Trigger TriggerType
	// Cron CRON表达式
	Cron string
	// Enabled 是否启用
	Enabled bool
	// DeletedTime 删除时间
	DeletedTime *time.Time
	// CreateTime 创建时间
	CreateTime time.Time
	// OwnerId 所有者id
	OwnerId int64
	// UpdateTime 修改时间
	UpdateTime time.Time
}

// LoadFromPO 将PO转换为实体
func LoadFromPO(po *taskPO.TaskPO) *Task {
	if po == nil {
		return nil
	}
	return &Task{
		TaskId:      po.TaskId,
		MaidId:      po.MaidId,
		TaskName:    po.TaskName,
		TaskType:    TaskType(po.TaskType),
		Trigger:     TriggerType(po.Trigger),
		Cron:        po.Cron,
		Enabled:     po.Enabled,
		DeletedTime: po.DeletedTime,
		CreateTime:  po.CreateTime,
		OwnerId:     po.OwnerId,
		UpdateTime:  po.UpdateTime,
	}
}
