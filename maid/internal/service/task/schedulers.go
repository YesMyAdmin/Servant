package task

import (
	taskEntity "common/public/model/entity/task"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/panjf2000/ants/v2"
)

const (
	POOL_SIZE = 8
)

//任务列表
var tasks map[string]*taskEntity.Task = make(map[string]*taskEntity.Task)

//从数据库加载任务列表
func LoadTasks() {

}

//从数据库获取任务详情,注册cron任务
func RegisterCronTask(taskId uint64) {
	cron := ""
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	scheduler := gocron.NewScheduler(timezone)
	scheduler.Cron(cron).Do(func() {
	})
}

//注销cron任务
func UnregisterCronTask(taskId uint64) {

}

var taskPool *TaskPool = nil


// TaskPool 任务执行池
type TaskPool struct{
	pool *ants.Pool
}

// 关闭任务池
func (pool *TaskPool) Destructor() {
	pool.pool.Release()
	pool = nil
}

// 提交一个异步任务
func (pool *TaskPool) Submit(task func()) {
	pool.pool.Submit(task)
}

// 获取一个任务池单例
func NewTaskPool() (*TaskPool, error) {
	if (nil != taskPool) {
		return taskPool, nil
	}
	pool, err := ants.NewPool(POOL_SIZE)
	if (err != nil) {
		return nil, err
	}
	taskPool = &TaskPool{pool: pool}
	return taskPool, nil
}