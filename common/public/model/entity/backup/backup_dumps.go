package backup

import (
	backupPO "common/internal/model/po/backup"
	"time"
)

// BackupDumpType 备份文件存储方式类型
type BackupDumpType string

const (
	// Local 本地文件系统
	Local BackupDumpType = "local"
	// Scp 异地主机(scp协议连接)
	Scp BackupDumpType = "scp"
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
	HoursToLive int32
	// Credential 认证信息,不同的存储方式有不同的认证信息
	Credential *string
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
		Credential:  po.Credential,
		Enabled:     po.Enabled,
		DeletedTime: po.DeletedTime,
		CreateTime:  po.CreateTime,
		OwnerId:     po.OwnerId,
		UpdateTime:  po.UpdateTime,
	}
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
	return &backupPO.BackupDumpPO{
		DumpId:      d.DumpId,
		DumpName:    d.DumpName,
		MaidId:      d.MaidId,
		DumpType:    string(d.DumpType),
		UrlTemplate: d.UrlTemplate,
		HoursToLive: d.HoursToLive,
		Credential:  d.Credential,
		Enabled:     d.Enabled,
		DeletedTime: d.DeletedTime,
		CreateTime:  d.CreateTime,
		OwnerId:     d.OwnerId,
		UpdateTime:  d.UpdateTime,
	}
}
