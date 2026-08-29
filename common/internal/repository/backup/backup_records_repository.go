package backup

import (
	"common/public/database"
	backupPO "common/internal/model/po/backup"
	"common/public/pkg"
	"time"
)

// ListBackupFiles 分页查询备份文件记录，支持按文件名模糊搜索
func ListBackupFiles(fileName string, pageNum, pageSize int) (*[]backupPO.BackupRecordPO, int64, error) {
	db := database.DB.Table("(?) AS ranked", 
			database.DB.Model(&backupPO.BackupRecordPO{}).
			Select("*, ROW_NUMBER() OVER (PARTITION BY file_id ORDER BY create_time DESC) AS rn")).
		Where("rn = 1") // 只保留每个文件的最新备份记录
	// 按名称模糊搜索
	if fileName != "" {
		db = db.Where("instr(file_name, ?)", fileName).Order("record_id DESC")
	}

	// 先查询总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, pkg.DatabaseError(backupPO.BackupRecordPO{}.TableName(), err)
	}

	// 分页查询
	var records []backupPO.BackupRecordPO
	offset := (pageNum - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Find(&records).Error; err != nil {
		return nil, 0, pkg.DatabaseError(backupPO.BackupRecordPO{}.TableName(), err)
	}

	return &records, total, nil
}


// ListRecordsByFileId 查询某个文件的备份记录
func ListRecordsByFileId(fileId uint64, pageNum, pageSize int) (*[]backupPO.BackupRecordPO, int64 ,error) {
	db := database.DB.Model(&backupPO.BackupRecordPO{}).Where("file_id = ?", fileId)
	var records []backupPO.BackupRecordPO

	var total int64
	countErr := db.Count(&total).Error
	if countErr != nil {
		return nil, 0, pkg.DatabaseError(backupPO.BackupRecordPO{}.TableName(), countErr)
	}

	offset := (pageNum - 1) * pageSize
	err := db.Offset(offset).Limit(pageSize).Find(&records).Error
	//返回数据库错误
	if err != nil {
		return nil, 0, pkg.DatabaseError(backupPO.BackupRecordPO{}.TableName(), err)
	}

	return &records, total, nil
}

// 
func SelectRecordsByFiles(files *[]uint64) (*[]backupPO.BackupRecordPO ,error) {
	if (len(*files) <= 0) {
		return nil, pkg.BadArgumentsError("files is empty");
	}
	if (len(*files) >= 200) {
		return nil, pkg.BadArgumentsError("files is too large");
	}
	var records []backupPO.BackupRecordPO
	dbError := database.DB.Model(&backupPO.BackupRecordPO{}).Where("file_id in (?)", files).Find(&records).Error
	//返回数据库错误
	if (dbError != nil) {
		return nil, pkg.DatabaseError(backupPO.BackupRecordPO{}.TableName(), dbError)
	}
	return &records, nil

}

// UpdateBackupFileId 将不同的文件id合并成一个
func UpdateBackupFileId(files *[]uint64, newFileId uint64) error {
	if (len(*files) <= 0) {
		return pkg.BadArgumentsError("files is empty");
	}
	if (len(*files) >= 200) {
		return pkg.BadArgumentsError("files is too large");
	}
	updateFileId := map[string]any{
    	"file_id": newFileId,
		"update_time": time.Now(),
	}
	dbError := database.DB.Model(&backupPO.BackupRecordPO{}).
		Where("file_id in (?)", files).
		Select("file_id").
		Updates(updateFileId).Error
	//返回数据库错误
	if (dbError != nil) {
		return pkg.DatabaseError(backupPO.BackupRecordPO{}.TableName(), dbError)
	}
	return nil
}

// DeleteBackupRecord 删除备份记录
func DeleteBackupRecord(backupRecordId uint64) error {
	db := database.DB.Model(&backupPO.BackupRecordPO{})
	err := db.Where("backup_record_id = ?", backupRecordId).Delete(&backupPO.BackupRecordPO{}).Error
	if err != nil {
		return pkg.DatabaseError(backupPO.BackupRecordPO{}.TableName(), err)
	}
	return nil
}