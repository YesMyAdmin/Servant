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
		Pages: (totalCount / req.PageSize) + 1,
		Contents: *responses,
	}, nil

	
}

// 查询备份文件详情
func BackupFileDetail(fileId string) {

}

// 手动合并文件的备份记录，这通常用于文件的路径改变后，管理员需要保持备份记录的连贯性
func MergeBackupFileRecords(files []string) {

}