package user

import "time"

type Core struct {
	ID          uint
	Name        string
	Email       string
	Password    string
	Address     string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserDataInterface interface {
	SelectAll() ([]Core, error)
	SelectById(id uint) (Core, error)
	Insert(input Core) error
	Login(email string, password string) (dataLogin Core, err error)
	Update(input Core) error
	Delete(input Core) error
}

type UserServiceInterface interface {
	GetAll() ([]Core, error)
	GetById(id uint) (Core, error)
	Create(input Core) error
	Login(email string, password string) (dataLogin Core, token string, err error)
	Update(input Core) error
	Delete(input Core) error
}
