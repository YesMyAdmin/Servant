package repository

import (
	"butler/internal/model/po"
	"butler/internal/database"
)	

// ListBackupFiles 分页查询备份文件记录，支持按文件名模糊搜索
func ListBackupFiles(fileName string, pageNum, pageSize int) (*[]po.BackupRecordPO, int64, error) {
	db := database.DB.Table("(?) AS ranked", 
			database.DB.Model(&po.BackupRecordPO{}).
			Select("*, ROW_NUMBER() OVER (PARTITION BY file_id ORDER BY create_time DESC) AS rn")).
		Where("rn = 1") // 只保留每个文件的最新备份记录
	// 按名称模糊搜索
	if fileName != "" {
		db = db.Where("instr(file_name, ?)", fileName).Order("record_id DESC")
	}

	// 先查询总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	var records []po.BackupRecordPO
	offset := (pageNum - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Find(&records).Error; err != nil {
		return nil, 0, err
	}

	return &records, total, nil
}


// ListRecordsByFileId 查询某个文件的备份记录
func ListRecordsByFileId(fileId uint64, pageNum, pageSize int) (*[]po.BackupRecordPO, int64 ,error) {
	db := database.DB.Model(&po.BackupRecordPO{}).Where("file_id = ?", fileId)
	var records []po.BackupRecordPO

	var total int64
	countErr := db.Count(&total).Error
	if countErr != nil {
		return nil, 0, countErr
	}

	offset := (pageNum - 1) * pageSize
	err := db.Offset(offset).Limit(pageSize).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return &records, total, nil
}

// DeleteBackupRecord 删除备份记录
func DeleteBackupRecord(backupRecordId uint64) error {
	db := database.DB.Model(&po.BackupRecordPO{})
	err := db.Where("backup_record_id = ?", backupRecordId).Delete(&po.BackupRecordPO{}).Error
	if err != nil {
		return err
	}
	return nil
}