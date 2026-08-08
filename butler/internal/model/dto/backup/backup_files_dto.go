package dto

import (
	"common/public/pkg"
	"time"
)

// BackupFileResp 备份文件查询响应
type BackupFileResp struct {
	FileId         string    `json:"fileId"`         // 文件ID
	FileName       string    `json:"fileName"`       // 文件名
	FileType       string    `json:"fileType"`       // 文件类型 file:文件 directory:目录 link:链接 other:其他(块设备、管道等)
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

// 备份记录查询响应
type BackupRecordResp struct {
	FileId         string    `json:"fileId"`         // 文件ID
	FileName       string    `json:"fileName"`       // 文件名
	FileType       string    `json:"fileType"`       // 文件类型 file:文件 directory:目录 link:链接 other:其他(块设备、管道等)
	MaidId         string    `json:"maidId"`         // 女仆节点ID
	TaskId         string    `json:"taskId"`         // 任务ID
	OriginalPath   string    `json:"originalPath"`   // 原始路径
	FileSize       string    `json:"fileSize"`       // 文件大小
	FileCreateTime time.Time `json:"fileCreateTime"` // 文件创建时间
	FileModifyTime time.Time `json:"fileModifyTime"` // 文件修改时间
	VersionHash    string    `json:"versionHash"`    // 版本哈希
	DumpId         string    `json:"dumpId"`         // 存储id
	DumpUrl        string    `json:"dumpUrl"`        // 存储位置url
	CreateTime     time.Time `json:"createTime"`     // 创建时间
	OwnerId        string    `json:"ownerId"`        // 所有者ID
	UpdateTime     time.Time `json:"updateTime"`     // 更新时间
}

// 列出最新备份文件
type ListBackupFileReq struct {
	FileName string `json:"fileName" form:"fileName"` //文件名称
	pkg.PageableReq
}

//列出某个文件的备份记录
type ListBackupFileRecordsReq struct {
	FileId string `json:"fileId" form:"fileId" binding:"required"` //文件id(string格式)
	pkg.PageableReq
}

// 手动合并文件备份记录请求
type MergeBackupFilesReq struct {
	Files []string `json:"files" binding:"required"` // 需要合并的文件ID列表
}

// 合并文件备份记录响应
type MergeBackupFilesResp struct {
	MergedFileId string `json:"mergedFileId"` // 合并后的文件ID（最新备份记录的文件ID）
}

// 删除备份记录请求
type DeleteBackupRecordReq struct {
	DeleteFile bool `json:"deleteFile"` // 是否同时删除存储的文件
}