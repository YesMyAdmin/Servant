package backup

import (
	backupRepo "common/internal/repository/backup"
	backupEntity "common/public/model/entity/backup"
)

// ListDumpsByTaskId 根据任务id查询存储方式列表
func ListDumpsByTaskId(taskId uint64) ([]*backupEntity.BackupDump, error) {
	dumpPOs, err := backupRepo.ListDumpsByTaskId(taskId)
	if err != nil {
		return nil, err
	}
	return backupEntity.LoadBackupDumpArray(dumpPOs), nil
}
