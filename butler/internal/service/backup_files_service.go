package service

import (
	"butler/internal/model/dto"
	"butler/internal/model/entity"
	"butler/internal/pkg"
	"butler/internal/repository"
)

// ListBackupFiles 列出备份文件
func ListBackupFiles(req *dto.ListBackupFileReq) (*pkg.PageableResp[dto.BackupFileResp], error) {
	backupFilesPO, totalCount, err := repository.ListBackupFiles(req.FileName, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	backupFiles := entity.LoadBackupFileRecordFromPOArray(backupFilesPO)
	responses := entity.DumpBackupFileRecordsToResp(backupFiles)
	return &pkg.PageableResp[dto.BackupFileResp]{
		Total: totalCount,
		Pages: int((totalCount / int64(req.PageSize)) + 1),
		Contents: *responses,
	}, nil

	
}

// 查询某个备份文件的备份记录
func BackupFileRecords(req *dto.ListBackupFileRecordsReq) (*pkg.PageableResp[dto.BackupRecordResp], error) {
	fileId, err := dto.StringToUint64(req.FileId)
	if (err != nil) {
		return nil, err
	}
	poArrays, total, repoErr := repository.ListRecordsByFileId(fileId, req.PageNum, req.PageSize)
	if (repoErr != nil) {
		return nil, repoErr
	}
	backupFiles := entity.LoadBackupFileRecordFromPOArray(poArrays)
	responses := entity.DumpBackupFileRecordsToRecordResp(backupFiles)
	return &pkg.PageableResp[dto.BackupRecordResp] {
		Total: total,
		Pages: int((total / int64(req.PageSize)) + 1),
		Contents: *responses,
	}, nil
	
}

// 手动合并文件的备份记录，这通常用于文件的路径改变后，管理员需要保持备份记录的连贯性
func MergeBackupFileRecords(files []string) {

}