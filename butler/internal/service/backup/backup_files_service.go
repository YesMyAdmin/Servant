package backup

import (
	dto "butler/internal/model/dto"
	backupdto "butler/internal/model/dto/backup"
	entity "butler/internal/model/entity/backup"
	backupRepo "butler/internal/repository/backup"
	"common/public/pkg"
)

// ListBackupFiles 列出备份文件
func ListBackupFiles(req *backupdto.ListBackupFileReq) (*pkg.PageableResp[backupdto.BackupFileResp], error) {
	backupFilesPO, totalCount, err := backupRepo.ListBackupFiles(req.FileName, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	backupFiles := entity.LoadBackupFileRecordFromPOArray(backupFilesPO)
	responses := entity.DumpBackupFileRecordsToResp(backupFiles)
	return &pkg.PageableResp[backupdto.BackupFileResp]{
		Total:    totalCount,
		Pages:    int((totalCount / int64(req.PageSize)) + 1),
		Contents: *responses,
	}, nil

}

// 查询某个备份文件的备份记录
func BackupFileRecords(req *backupdto.ListBackupFileRecordsReq) (*pkg.PageableResp[backupdto.BackupRecordResp], error) {
	fileId, err := dto.StringToUint64(req.FileId)
	if err != nil {
		return nil, err
	}
	poArrays, total, repoErr := backupRepo.ListRecordsByFileId(fileId, req.PageNum, req.PageSize)
	if repoErr != nil {
		return nil, repoErr
	}
	backupFiles := entity.LoadBackupFileRecordFromPOArray(poArrays)
	responses := entity.DumpBackupFileRecordsToRecordResp(backupFiles)
	return &pkg.PageableResp[backupdto.BackupRecordResp]{
		Total:    total,
		Pages:    int((total / int64(req.PageSize)) + 1),
		Contents: *responses,
	}, nil

}

// 手动合并文件的备份记录，这通常用于文件的路径改变后，管理员需要保持备份记录的连贯性
func MergeBackupFileRecords(files *[]uint64) (uint64, error) {
	recordsPO, err := backupRepo.SelectRecordsByFiles(files)
	if err != nil {
		return 0, err
	}
	records := entity.LoadBackupFileRecordFromPOArray(recordsPO)
	//将多个文件并入到最新记录对应的文件中
	var newestRecord *entity.BackupFileRecord
	var fileType string
	for _, record := range *records {
		if newestRecord == nil || newestRecord.CreateTime.Before(record.CreateTime) {
			newestRecord = &record
		}
		//所有合并文件的类型必须一致
		if fileType == "" {
			fileType = record.FileType
		} else if fileType != record.FileType {
			return 0, pkg.FileMergingConflictError(files)
		}
	}
	//更新旧有记录的文件id
	dbErr := backupRepo.UpdateBackupFileId(files, newestRecord.BackupRecordId)
	if dbErr != nil {
		return 0, dbErr
	}
	return newestRecord.BackupRecordId, nil
}

// DeleteBackupRecord 删除备份记录
func DeleteBackupRecord(backupRecordId uint64, deleteFile bool) error {
	// 如果 deleteFile 为 true，还需要实际删除存储的文件
	if deleteFile {
		// TODO: 实现文件删除逻辑
	}
	return backupRepo.DeleteBackupRecord(backupRecordId)
}
