package task

import "time"

type Core struct {
	ID         uint
	Name       string
	ProjectID  uint
	DetailTask string
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Project    ProjectCore
}

type ProjectCore struct {
	ID            uint
	Name          string
	DetailProject string
}

type TaskDataInterface interface {
	Insert(input Core) error
	// SelectAll() ([]Core, error)
	// SelectById(id uint) (Core, error)
	Update(id uint, input Core) error
	Delete(id uint) error
}

type TaskServiceInterface interface {
	Create(input Core) error
	// GetAll() ([]Core, error)
	// GetById(id uint) (Core, error)
	Update(id uint, input Core) error
	Deletes(id uint) error
}
