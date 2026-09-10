package task

import (
	taskRepo "common/internal/repository/task"
	taskDto "common/public/model/dto/task"
	taskEntity "common/public/model/entity/task"
)

// GetTask 根据任务ID查询定时任务
func GetTask(taskId uint64) (*taskEntity.Task, error) {
	po, err := taskRepo.SelectById(taskId)
	if err != nil {
		return nil, err
	}
	return taskEntity.LoadFromPO(po), nil
}

// ListTasks 根据可选条件查询定时任务
func ListTasks(req *taskDto.ListTasksReq) ([]taskEntity.Task, error) {
	pos, err := taskRepo.ListTasks(req.MaidId, req.TaskName, req.TaskType)
	if err != nil {
		return nil, err
	}

	tasks := make([]taskEntity.Task, 0, len(pos))
	for i := range pos {
		if t := taskEntity.LoadFromPO(&pos[i]); t != nil {
			tasks = append(tasks, *t)
		}
	}
	return tasks, nil
}
