package backup

// BackupTaskDumpPO 备份任务和存储方式关联表 PO (对应表 backup_task_dumps)
type BackupTaskDumpPO struct {
	TaskId uint64 `gorm:"column:task_id;primaryKey;comment:任务id" json:"taskId"`
	DumpId uint64 `gorm:"column:dump_id;primaryKey;comment:存储方式id" json:"dumpId"`
}

// TableName 指定表名
func (BackupTaskDumpPO) TableName() string {
	return "backup_task_dumps"
}
