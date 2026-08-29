package task

import "time"

// TaskPO 定时任务表 PO (对应表 tasks)
type TaskPO struct {
	TaskId      uint64     `gorm:"column:task_id;primaryKey;comment:任务ID" json:"taskId"`
	MaidId      uint64     `gorm:"column:maid_id;not null;comment:女仆节点id" json:"maidId"`
	TaskName    string     `gorm:"column:task_name;type:varchar(32);not null;comment:定时任务名称" json:"taskName"`
	TaskType    string     `gorm:"column:task_type;type:varchar(12);not null;comment:任务类型(builtin:系统内置的任务 backup:备份任务)" json:"taskType"`
	Trigger     string     `gorm:"column:trigger;type:varchar(12);not null;comment:触发器类型(cron:cron表达式)" json:"trigger"`
	Cron        string     `gorm:"column:cron;type:varchar(32);comment:CRON表达式" json:"cron"`
	Enabled     bool       `gorm:"column:enabled;type:tinyint(1);not null;default:1;comment:是否启用(1:是 0:否)" json:"enabled"`
	DeletedTime *time.Time `gorm:"column:deleted_time;default:null;comment:删除时间(留空为未删除,有时间为软删除)" json:"deletedTime"`
	CreateTime  time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:记录创建时间" json:"createTime"`
	OwnerId     int64      `gorm:"column:owner_id;not null;comment:所有者id(默认创建者)" json:"ownerId"`
	UpdateTime  time.Time  `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:记录修改时间" json:"updateTime"`
}

// TableName 指定表名
func (TaskPO) TableName() string {
	return "tasks"
}
