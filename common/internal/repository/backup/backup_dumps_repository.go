package backup

import (
	"common/public/database"
	backupPO "common/internal/model/po/backup"
	"time"
)

// ListDumps 分页查询备份存储方式，支持按名称模糊搜索
func ListDumps(pageNum, pageSize int, dumpName string) ([]backupPO.BackupDumpPO, int64, error) {
	db := database.DB.Model(&backupPO.BackupDumpPO{})

	// 按名称模糊搜索
	if dumpName != "" {
		db = db.Where("instr(dump_name, ?)", dumpName)
	}

	// 先查询总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	var dumps []backupPO.BackupDumpPO
	offset := (pageNum - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Find(&dumps).Error; err != nil {
		return nil, 0, err
	}

	return dumps, total, nil
}

// ListDumpsByTaskId 根据任务id查询其关联的存储方式列表
func ListDumpsByTaskId(taskId uint64) ([]backupPO.BackupDumpPO, error) {
	dumpTable := backupPO.BackupDumpPO{}.TableName()
	relationTable := backupPO.BackupTaskDumpPO{}.TableName()

	var dumps []backupPO.BackupDumpPO
	err := database.DB.Model(&backupPO.BackupDumpPO{}).
		Joins("JOIN "+relationTable+" ON "+relationTable+".dump_id = "+dumpTable+".dump_id").
		Where(relationTable+".task_id = ?", taskId).
		Find(&dumps).Error
	if err != nil {
		return nil, err
	}
	return dumps, nil
}

// SelectBackupDump 查询备份存储方式
func SelectBackupDump(dumpId uint64) (*backupPO.BackupDumpPO, error) {
	db := database.DB.Model(&backupPO.BackupDumpPO{})
	var dump backupPO.BackupDumpPO
	err := db.Where("dump_id = ?", dumpId).First(&dump).Error
	if err != nil {
		return nil, err
	}
	return &dump, nil
}

// NewBackupDump 添加新的备份存储方式
func NewBackupDump(backupDump *backupPO.BackupDumpPO) error {
	db := database.DB.Model(&backupPO.BackupDumpPO{})
	err := db.Create(backupDump).Error
	if err != nil {
		return err
	}
	return nil
}

// EditBackupDump 更新备份存储方式
func EditBackupDump(backupDump *backupPO.BackupDumpPO) error {
	db := database.DB.Model(&backupPO.BackupDumpPO{})
	err := db.Where("dump_id = ?", backupDump.DumpId).Updates(backupDump).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteBackupDump 删除备份存储方式(软删除)
func DeleteBackupDump(dumpId uint64) error {
	db := database.DB.Model(&backupPO.BackupDumpPO{})
	err := db.Where("dump_id = ?", dumpId).UpdateColumn("deleted_time", time.Now()).Error
	if err != nil {
		return err
	}
	return nil
}
