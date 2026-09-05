package task

import "github.com/panjf2000/ants/v2"

const (
	POOL_SIZE = 8
)

var taskPool *TaskPool = nil


// TaskPool 任务池
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