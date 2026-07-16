package dto

import "time"

// BackupFileResp 备份文件查询响应
type BackupFileResp struct {
	FileId         string    `json:"fileId"`         // 文件ID
	FileName       string    `json:"fileName"`       // 文件名
	MaidId         string    `json:"maidId"`         // 女仆节点ID
	TaskId         string    `json:"taskId"`         // 任务ID
	OriginalPath   string    `json:"originalPath"`   // 原始路径
	FileSize       string    `json:"fileSize"`       // 文件大小
	FileCreateTime time.Time `json:"fileCreateTime"` // 文件创建时间
	FileModifyTime time.Time `json:"fileModifyTime"` // 文件修改时间
	VersionHash    string    `json:"versionHash"`    // 版本哈希
	CreateTime     time.Time `json:"createTime"`     // 创建时间
	OwnerId        string    `json:"ownerId"`        // 所有者ID
	UpdateTime     time.Time `json:"updateTime"`     // 更新时间
}