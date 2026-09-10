package dto

// ListTasksReq 查询定时任务的请求条件
// 各条件均为可选, 不传即不按该条件过滤
type ListTasksReq struct {
	// MaidId 女仆节点id, 传 0 表示不过滤
	MaidId uint64 `json:"maidId" form:"maidId"`
	// TaskName 定时任务名称(模糊匹配), 传空字符串表示不过滤
	TaskName string `json:"taskName" form:"taskName"`
	// TaskType 任务类型(builtin:系统内置的任务 backup:备份任务), 传空字符串表示不过滤
	TaskType string `json:"taskType" form:"taskType"`
}
