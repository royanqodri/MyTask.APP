package handler

import (
	"royan/cleanarch/features/task"
)

type TaskRequest struct {
	Name       string `json:"name" form:"name"`
	ProjectID  uint   `json:"project_id" form:"project_id"`
	DetailTask string `json:"detail_task" form:"detail_task"`
	Status     string `json:"status" form:"status"`
}

func RequestToCore(input TaskRequest) task.Core {
	return task.Core{
		Name:       input.Name,
		DetailTask: input.DetailTask,
		Status:     input.Status,
	}
}
