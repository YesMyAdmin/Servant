package model

type Tasks struct {
	TaskId      uint64  
	MaidId 		uint64
}

// BackupTask 备份任务
type BackupTask struct {
	TaskId      uint64  
	MaidId 		uint64
	Mode        string    
	Source      string   
	DeletedTime time.Time 
	CreateTime  time.Time 
	OwnerId     int64     
	UpdateTime  time.Time 
}