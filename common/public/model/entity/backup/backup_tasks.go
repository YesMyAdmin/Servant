package backup

import (
	"common/public/model/dto"
	backupdto "common/public/model/dto/backup"
	backupPO "common/internal/model/po/backup"
)

// ToPO 将 NewBackupTaskReq 转换为 BackupTaskPO
// 仅映射 PO 中存在的字段：Mode、Source
func NewReqToPO(r *backupdto.NewBackupTaskReq) *backupPO.BackupTaskPO {
	return &backupPO.BackupTaskPO{
		Mode:   string(r.Mode),
		Source: r.Source,
	}
}

// ToPO 将 EditBackupTaskReq 转换为 BackupTaskPO
// 仅映射 PO 中存在的字段：TaskId、Mode、Source
func EditReqToPO(r *backupdto.EditBackupTaskReq) *backupPO.BackupTaskPO {
	taskId, err := dto.StringToUint64(r.TaskId)
	if err != nil {
		return nil
	}
	return &backupPO.BackupTaskPO{
		TaskId: taskId,
		Mode:   string(r.Mode),
		Source: r.Source,
	}
}

// ToListTasksResp 将 BackupTaskPO 转换为 ListTasksResp
// 映射 PO 中所有公共字段；MaidId、MaidName、Cron、Enabled 等
// 不属于 PO 的字段由调用方按需填充
func ToListTasksResp(p *backupPO.BackupTaskPO) *backupdto.ListTasksResp {
	if p == nil {
		return nil
	}
	return &backupdto.ListTasksResp{
		TaskId:     dto.Uint64ToString(p.TaskId),
		Mode:       p.Mode,
		Source:     p.Source,
		CreateTime: p.CreateTime,
		OwnerId:    dto.Uint64ToString(uint64(p.OwnerId)),
		UpdateTime: p.UpdateTime,
	}
}

// ToListTasksRespSlice 批量将 BackupTaskPO 切片转换为 ListTasksResp 切片
func ToListTasksRespSlice(pos []backupPO.BackupTaskPO) []backupdto.ListTasksResp {
	if pos == nil {
		return nil
	}
	result := make([]backupdto.ListTasksResp, 0, len(pos))
	for i := range pos {
		if resp := ToListTasksResp(&pos[i]); resp != nil {
			result = append(result, *resp)
		}
	}
	return result
}