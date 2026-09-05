package backup

import (
	backupPO "common/internal/model/po/backup"
	"encoding/json"
	"time"
)

// BackupDumpType 备份文件存储方式类型
type BackupDumpType string

const (
	// Local 本地文件系统
	Local BackupDumpType = "local"
	// SCP 异地主机(scp协议连接)
	SCP BackupDumpType = "scp"
	// S3 对象存储s3协议
	S3 BackupDumpType = "s3"
)

// BackupDump 备份文件存储方式实体
type BackupDump struct {
	// DumpId 存储位置id
	DumpId uint64
	// DumpName 存储方式名称
	DumpName string
	// MaidId 女仆节点id
	MaidId uint64
	// DumpType 存储类型
	DumpType BackupDumpType
	// UrlTemplate 存储位置url模板
	UrlTemplate string
	// HoursToLive 保存时间(以小时为单位),过期由定时任务清理文件
	HoursToLive *int32
	// Credential 认证信息,不同的存储方式有不同的认证信息
	Credential *DumpCredential
	// Enabled 是否启用
	Enabled bool
	// DeletedTime 删除时间
	DeletedTime *time.Time
	// CreateTime 创建时间
	CreateTime time.Time
	// OwnerId 所有者id
	OwnerId uint64
	// UpdateTime 修改时间
	UpdateTime time.Time
}

// LoadBackupDump 将 BackupDumpPO 转换为 BackupDump 实体
func LoadBackupDump(po *backupPO.BackupDumpPO) *BackupDump {
	if po == nil {
		return nil
	}
	return &BackupDump{
		DumpId:      po.DumpId,
		DumpName:    po.DumpName,
		MaidId:      po.MaidId,
		DumpType:    BackupDumpType(po.DumpType),
		UrlTemplate: po.UrlTemplate,
		HoursToLive: po.HoursToLive,
		Credential:  loadCredential(BackupDumpType(po.DumpType), po.Credential),
		Enabled:     po.Enabled,
		DeletedTime: po.DeletedTime,
		CreateTime:  po.CreateTime,
		OwnerId:     po.OwnerId,
		UpdateTime:  po.UpdateTime,
	}
}

// loadCredential 根据存储类型将 JSON 字符串解析为对应的认证信息
func loadCredential(dumpType BackupDumpType, raw *string) *DumpCredential {
	if raw == nil {
		return nil
	}
	var credential DumpCredential
	switch dumpType {
	case SCP:
		credential = &SCPDumpCredential{}
	case S3:
		credential = &S3DumpCredential{}
	default:
		credential = &LocalDumpCredential{}
	}
	if jsonError := json.Unmarshal([]byte(*raw), credential); jsonError != nil {
		return nil
	}
	return &credential
}

// LoadBackupDumpArray 将 BackupDumpPO 切片转换为 BackupDump 实体切片
func LoadBackupDumpArray(poArray []backupPO.BackupDumpPO) []*BackupDump {
	dumps := make([]*BackupDump, 0, len(poArray))
	for i := range poArray {
		if dump := LoadBackupDump(&poArray[i]); dump != nil {
			dumps = append(dumps, dump)
		}
	}
	return dumps
}

// ToPO 将 BackupDump 实体转换为 BackupDumpPO
func (d *BackupDump) ToPO() *backupPO.BackupDumpPO {
	if d == nil {
		return nil
	}
	bytes, jsonError := json.Marshal(d.Credential)
	var credential string
	if (jsonError != nil) {
		credential = ""
	} else {
		//转为UTF8字符串
		credential = string(bytes)
	}
	return &backupPO.BackupDumpPO{
		DumpId:      d.DumpId,
		DumpName:    d.DumpName,
		MaidId:      d.MaidId,
		DumpType:    string(d.DumpType),
		UrlTemplate: d.UrlTemplate,
		HoursToLive: d.HoursToLive,
		Credential:  &credential,
		Enabled:     d.Enabled,
		DeletedTime: d.DeletedTime,
		CreateTime:  d.CreateTime,
		OwnerId:     d.OwnerId,
		UpdateTime:  d.UpdateTime,
	}
}

// 转储目标认证方式
type DumpCredential interface {
	DumpType() BackupDumpType
}

// 本地转储认证方式(其实是空的)
type LocalDumpCredential struct {
}

func (c *LocalDumpCredential) DumpType() BackupDumpType {
	return Local
}

type SCPDumpCredential struct {
	// 用户名(不使用root,如果输入"root",前后端报错拒绝)
	User string `json:"user"`
	// ssh登录密钥/证书,支持填写密钥原文和文件路径
	// 若使用原文,则数据库端加密,不返回给前端
	Secret string `json:"secret"`
}

func (c *SCPDumpCredential) DumpType() BackupDumpType {
	return SCP
}

type S3DumpCredential struct {
	// keyId
	AccessKeyId string `json:"accessKeyId"`
	// 访问密钥,数据库端加密,不返回给前端
	AccessKey string `json:"accessKey"`
}

func (c *S3DumpCredential) DumpType() BackupDumpType {
	return S3
}
