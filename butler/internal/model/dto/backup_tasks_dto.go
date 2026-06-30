package dto

import (
	"butler/internal/pkg"
	"time"
)

// 新建备份任务请求
type NewBackupTaskReq struct {
	// MaidId 负责执行任务的女仆节点ID
	MaidId int64 `json:"maidId" binding:"required"`

	// TaskName 备份任务名称
	TaskName string `json:"taskName" binding:"required"`

	// Mode 备份模式: full-全量, incremental-增量
	Mode BackupMode `json:"mode" binding:"required,oneof=full incremental"`

	// 执行备份的cron表达式
	Cron string `json:"cron" binding:"required"`

	// Source 需要备份的文件/文件夹路径
	Source string `json:"source" binding:"required"`

	// Enabled 任务开关，默认false(关闭)
	Enabled bool `json:"enabled" default:"false"`

	// Dumps 文件存储方式ID列表
	Dumps []int64 `json:"dumps" binding:"required,min=1"`
}

// BackupMode 备份模式枚举
type BackupMode string

const (
	// BackupModeFull 全量备份
	BackupModeFull BackupMode = "full"

	// BackupModeIncremental 增量备份
	BackupModeIncremental BackupMode = "incremental"
)

// 查询定时任务的请求
type ListTasksReq struct {
	TaskName string `json:"taskName" form:"taskName"` //任务名称
	pkg.PageableReq
}

// 查询定时任务的返回体
type ListTasksResp struct {
	TaskId         	uint64     	`json:"taskId"`     // 任务ID
	MaidId         	uint64	    `json:"maidId"`     // 女仆节点ID
	MaidName       	string    	`json:"maidName"`   // 女仆节点名称
	Mode 			string 		`json:"mode"`		// 备份模式: full-全量, incremental-增量
	Source   	   	string    	`json:"source"`		// 备份路径
	Cron 			string 		`json:"cron"`		// CRON表达式
	Enabled 		bool 		`json:"enabled"`	// 是否启用
	CreateTime 		time.Time 	`json:"createTime"`	// 记录创建时间
	OwnerId 		uint64 		`json:"ownerId"`	// 所有者id
	UpdateTime 		time.Time 	`json:"updateTime"`	// 记录修改时间
}
