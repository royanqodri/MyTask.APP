package project

import "time"

type Core struct {
	ID            uint
	Name          string
	UserID        uint
	DetailProject string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	User          UserCore
}

type UserCore struct {
	ID    uint
	Name  string
	Email string
}

type ProjectDataInterface interface {
	Insert(input Core) error
	SelectAll() ([]Core, error)
	SelectById(id uint) (Core, error)
	Update(id uint, input Core) error
	Delete(id uint) error
}

type ProjectServiceInterface interface {
	Create(input Core) error
	GetAll() ([]Core, error)
	GetById(id uint) (Core, error)
	Update(id uint, input Core) error
	Deletes(id uint) error
}
