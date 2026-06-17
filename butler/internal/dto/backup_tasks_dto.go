package dto

//查询定时任务的请求
type ListTasksReq struct {
	TaskName string
	PageableReq
}

