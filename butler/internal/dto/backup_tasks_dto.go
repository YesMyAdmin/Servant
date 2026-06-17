package dto

import (
	"butler/internal/pkg"
	"time"
)

//查询定时任务的请求
type ListTasksReq struct {
	TaskName string `json:"taskName" form:"taskName"` //任务名称
	pkg.PageableReq
}

//查询定时任务的返回体
type ListTasksResp struct {
	FileId         int64     `json:"fileId"`         // 文件ID
	MaidId         int64     `json:"maidId"`         // 女仆ID
	MaidName       string    `json:"maidName"`       // 女仆名称
	OriginalPath   string    `json:"originalPath"`   // 原始路径
	FileSize       uint64    `json:"fileSize"`       // 文件大小
	FileModifyTime time.Time `json:"fileModifyTime"` // 文件修改时间
	VersionHash    uint64    `json:"versionHash"`    // 版本哈希 (uint64 避免溢出)
	RecordTime     time.Time `json:"recordTime"`     // 记录时间
	UpdateTime     time.Time `json:"updateTime"`     // 更新时间
}

