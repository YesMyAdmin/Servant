package entity

import (
	"bufio"
	"butler/internal/model/dto"
	"butler/internal/model/po"
	"io"
	"os"
	"time"
	"github.com/cespare/xxhash/v2"
)

// BackupFileRecord 备份文件记录实体
type BackupFileRecord struct {
	// BackupRecordId 备份记录ID
	BackupRecordId        uint64
	// FileId 文件ID
	FileId                uint64
	// FileName 文件名
	FileName              string
	// MaidId 女仆节点ID
	MaidId                uint64
	// TaskId 任务ID
	TaskId                uint64
	// OriginalPath 原始路径
	OriginalPath          string
	// FileSize 文件大小
	FileSize              uint64
	// FileCreateTime 文件创建时间
	FileCreateTime        time.Time
	// FileModifyTime 文件修改时间
	FileModifyTime        time.Time
	// VersionHash 版本哈希
	VersionHash           uint64
	// DumpId 转储ID
	DumpId                uint64
	// ExpireTime 过期时间
	ExpireTime            time.Time
	// FileActualDeletedTime 文件实际删除时间
	FileActualDeletedTime time.Time
	// DeletedTime 删除时间
	DeletedTime           time.Time
	// CreateTime 创建时间
	CreateTime            time.Time
	// OwnerId 所有者ID
	OwnerId               uint64
	// UpdateTime 更新时间
	UpdateTime            time.Time
}

// 计算文件id
func (b *BackupFileRecord) computeFileId() {
	maidIdBytes := make([]byte, 8)
	for i := range 8 {
		maidIdBytes[i] = byte((b.MaidId >> (56 - 8*i)) & 0xFF)
	}
	originalPathBytes := []byte(b.OriginalPath)
	combined := append(maidIdBytes, originalPathBytes...)
	b.FileId = xxhash.Sum64(combined)
}

// 计算文件版本哈希
func (b *BackupFileRecord) computeVersionHash() error {
	file, err := os.Open(b.OriginalPath)
	if err != nil {
		defer file.Close()
		return err
	}
	reader := bufio.NewReader(file)
	readBuffer := make([]byte, 1024 * 1024 * 16) // 16MB 缓冲区
	versionHashs := xxhash.New()
	for {
		bytesOfRead, err := reader.Read(readBuffer)
		if err == io.EOF || bytesOfRead == 0{
			break
		}
		if err != nil {
			defer file.Close()
			return err
		}
		versionHashs.Write(readBuffer[:bytesOfRead])
	}
	b.VersionHash = versionHashs.Sum64()
	defer file.Close()
	return nil
}

// DumpToPO 将 BackupFile 转换为 BackupRecordPO
func (b *BackupFileRecord) DumpToPO() *po.BackupRecordPO {
	return &po.BackupRecordPO{
		BackupRecordId:        b.BackupRecordId,
		FileId:                b.FileId,
		FileName:              b.FileName,
		MaidId:                b.MaidId,
		TaskId:                b.TaskId,
		OriginalPath:          b.OriginalPath,
		FileSize:              b.FileSize,
		FileCreateTime:        b.FileCreateTime,
		FileModifyTime:        b.FileModifyTime,
		VersionHash:           b.VersionHash,
		DumpId:                b.DumpId,
		ExpireTime:            b.ExpireTime,
		FileActualDeletedTime: b.FileActualDeletedTime,
		DeletedTime:           b.DeletedTime,
		CreateTime:            b.CreateTime,
		OwnerId:               b.OwnerId,
		UpdateTime:            b.UpdateTime,
	}
}

func (b *BackupFileRecord) DumpToResp() *dto.BackupFileResp {
	return &dto.BackupFileResp{
		FileId:         dto.Uint64ToString(b.FileId),
		FileName:       b.FileName,
		MaidId:         dto.Uint64ToString(b.MaidId),
		TaskId:         dto.Uint64ToString(b.TaskId),
		OriginalPath:   b.OriginalPath,
		FileSize:       dto.Uint64ToString(b.FileSize),
		FileCreateTime: b.FileCreateTime,
		FileModifyTime: b.FileModifyTime,
		VersionHash:    dto.Uint64ToString(b.VersionHash),
		CreateTime:     b.CreateTime,
		OwnerId:        dto.Uint64ToString(b.OwnerId),
		UpdateTime:     b.UpdateTime,
	}
}

// LoadBackupFileRecordFromPO 将 BackupRecordPO 转换为 BackupFile 实体
func LoadBackupFileRecordFromPO(po *po.BackupRecordPO) *BackupFileRecord {
	return &BackupFileRecord{
		BackupRecordId:        po.BackupRecordId,
		FileId:                po.FileId,
		FileName:              po.FileName,
		MaidId:                po.MaidId,
		TaskId:                po.TaskId,
		OriginalPath:          po.OriginalPath,
		FileSize:              po.FileSize,
		FileCreateTime:        po.FileCreateTime,
		FileModifyTime:        po.FileModifyTime,
		VersionHash:           po.VersionHash,
		DumpId:                po.DumpId,
		ExpireTime:            po.ExpireTime,
		FileActualDeletedTime: po.FileActualDeletedTime,
		DeletedTime:           po.DeletedTime,
		CreateTime:            po.CreateTime,
		OwnerId:               po.OwnerId,
		UpdateTime:            po.UpdateTime,
	}
}

// 将BackupRecordPO数组转换为BackupFileRecord数组
func LoadBackupFileRecordFromPOArray(poArray *[]po.BackupRecordPO) *[]BackupFileRecord {
	var entityArray []BackupFileRecord
	for _, po := range *poArray {
		entityArray = append(entityArray, *LoadBackupFileRecordFromPO(&po))
	}
	return &entityArray
}

// 将BackupFileRecord数组转换为BackupFileResp返回体数组
func DumpBackupFileRecordsToResp(entityArray *[]BackupFileRecord) *[]dto.BackupFileResp {
	var respArray []dto.BackupFileResp
	for _, entity := range *entityArray {
		respArray = append(respArray, *entity.DumpToResp())
	}
	return &respArray
}