package task

import (
	taskEntity "common/public/model/entity/task"
	taskService "common/public/service/task"
	backupTask "maid/internal/service/task/backup"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
)

type TaskExecutor struct {
	Task taskEntity.Task
	Runner func(uint64)
}

func (t *TaskExecutor) Run() {
	t.Runner(t.Task.TaskId)
}

var scheduler = initScheduler()

func initScheduler() (*gocron.Scheduler){
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	return gocron.NewScheduler(timezone)
}

func GetTaskExecutor(taskId uint64) (*TaskExecutor, error) {
	task, err := taskService.GetTask(taskId)
	if (err != nil) {
		return nil, err
	}
	switch(task.TaskType){
	case taskEntity.Backup:
		return &TaskExecutor{
			Task: *task,
			Runner: func(taskId uint64) {
				backupTask.DoBackup(taskId)
			},
	}, nil
	}
	return nil, nil
}

//从数据库获取任务详情,注册cron任务
func RegisterCronTask(taskId uint64) error {
	cron := ""
	// timezone, _ := time.LoadLocation("Asia/Shanghai")
	// scheduler := gocron.NewScheduler(timezone)
	executor, err := GetTaskExecutor(taskId)
	if (err != nil) {
		return err
	}
	job, _ := scheduler.Cron(cron).Do(executor.Run)
	job.Tag(strconv.FormatUint(taskId, 10))
	return nil
}

//注销cron任务
func UnregisterCronTask(taskId uint64) {
	scheduler.RemoveByTag(strconv.FormatUint(taskId, 10))
}