package service

import (
	"errors"
	"royan/cleanarch/features/task"
)

type taskService struct {
	taskData task.TaskDataInterface
}

// Update implements task.TaskServiceInterface
func (service *taskService) Update(id uint, input task.Core) error {
	err := service.taskData.Update(id, input)
	return err
}

// Deletes implements task.TaskServiceInterface
func (service *taskService) Deletes(id uint) error {
	err := service.taskData.Delete(id)
	return err
}

// Create implements task.TaskServiceInterface
func (service *taskService) Create(input task.Core) error {
	if input.Name == "" || input.DetailTask == "" || input.Status == "" {
		return errors.New("validation error. name/detail/status required")
	}
	err := service.taskData.Insert(input)
	return err
}

func New(repo task.TaskDataInterface) task.TaskServiceInterface {
	return &taskService{
		taskData: repo,
	}
}
