package service

import (
	"errors"
	"royan/cleanarch/app/middlewares"
	"royan/cleanarch/features/user"
)

type userService struct {
	userData user.UserDataInterface
}

// Delete implements user.UserServiceInterface
func (service *userService) Delete(input user.Core) error {
	if input.Email == "" && input.Password == "" {
		return errors.New("validasi error. at least one field (email/password) is required for update")
	}

	// Memanggil metode Update dari UserDataInterface
	err := service.userData.Update(input)
	return err
}

// Update implements user.UserServiceInterface
func (service *userService) Update(input user.Core) error {
	if input.Name == "" && input.Email == "" && input.Password == "" {
		return errors.New("validasi error. at least one field (name/email/password) is required for update")
	}

	// Memanggil metode Update dari UserDataInterface
	err := service.userData.Update(input)
	return err
}

// GetAll implements user.UserServiceInterface
func (service *userService) GetAll() ([]user.Core, error) {
	result, err := service.userData.SelectAll()
	return result, err
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
	}
}

// Login implements user.UserServiceInterface.
func (service *userService) Login(email string, password string) (dataLogin user.Core, token string, err error) {
	dataLogin, err = service.userData.Login(email, password)
	if err != nil {
		return user.Core{}, "", err
	}
	token, err = middlewares.CreateToken(int(dataLogin.ID))
	if err != nil {
		return user.Core{}, "", err
	}
	return dataLogin, token, nil
}

// GetById implements user.UserServiceInterface.
func (service *userService) GetById(id uint) (user.Core, error) {
	return service.userData.SelectById(id)
}

// Create implements user.UserServiceInterface.
func (service *userService) Create(input user.Core) error {
	if input.Name == "" || input.Email == "" || input.Password == "" {
		return errors.New("validation error. name/email/password required")
	}
	err := service.userData.Insert(input)
	return err
}
