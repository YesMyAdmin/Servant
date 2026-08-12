package backup

import (
	"butler/internal/model/dto/backup"
	backupsvc "butler/internal/service/backup"
	"common/public/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListBackupFiles 列出最新备份文件
func ListBackupFiles(c *gin.Context) error {
	var req dto.ListBackupFileReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		return pkg.BadArgumentsError(err.Error())
	}
	resp, svcErr := backupsvc.ListBackupFiles(&req)
	if svcErr != nil {
		return svcErr
	}
	c.JSON(http.StatusOK, resp)
	return nil
}

// ListBackupFileRecords 查询某个备份文件的备份记录
func ListBackupFileRecords(c *gin.Context) error {
	var req dto.ListBackupFileRecordsReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		return pkg.BadArgumentsError(err.Error())
	}
	resp, svcErr := backupsvc.BackupFileRecords(&req)
	if svcErr != nil {
		return svcErr
	}
	c.JSON(http.StatusOK, resp)
	return nil
}

// MergeBackupFileRecords 手动合并文件的备份记录
func MergeBackupFileRecords(c *gin.Context) error {
	var req dto.MergeBackupFilesReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return pkg.BadArgumentsError(err.Error())
	}

	// 转换 []string 为 []uint64
	fileIds := make([]uint64, 0, len(req.Files))
	for _, f := range req.Files {
		id, parseErr := strconv.ParseUint(f, 10, 64)
		if parseErr != nil {
			return pkg.BadArgumentsError("invalid file id: " + f)
		}
		fileIds = append(fileIds, id)
	}

	mergedFileId, svcErr := backupsvc.MergeBackupFileRecords(&fileIds)
	if svcErr != nil {
		return svcErr
	}

	resp := dto.MergeBackupFilesResp{
		MergedFileId: strconv.FormatUint(mergedFileId, 10),
	}
	c.JSON(http.StatusOK, resp)
	return nil
}

// DeleteBackupRecord 删除备份记录
func DeleteBackupRecord(c *gin.Context) error {
	backupRecordIdStr := c.Request.PathValue("backupRecordId")
	backupRecordId, parseErr := strconv.ParseUint(backupRecordIdStr, 10, 64)
	if parseErr != nil {
		return pkg.BadArgumentsError(parseErr.Error())
	}

	var req dto.DeleteBackupRecordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return pkg.BadArgumentsError(err.Error())
	}

	svcErr := backupsvc.DeleteBackupRecord(backupRecordId, req.DeleteFile)
	if svcErr != nil {
		return svcErr
	}

	c.JSON(http.StatusOK, pkg.SuccessMessageResp(""))
	return nil
}
