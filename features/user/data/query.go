package data

import (
	"errors"
	"royan/cleanarch/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// Delete implements user.UserDataInterface
func (repo *userQuery) Delete(input user.Core) error {
	var userGorm User
	tx := repo.db.First(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}

	// Hapus pengguna dari database
	tx = repo.db.Delete(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("data not found to deleted")

	}
	return nil
}

// Update implements user.UserDataInterface
func (repo *userQuery) Update(input user.Core) error {
	var userGorm User
	tx := repo.db.First(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}

	// Perbarui properti pengguna dengan data dari input
	userGorm.Name = input.Name
	userGorm.Email = input.Email
	userGorm.Password = input.Password
	userGorm.Address = input.Address
	userGorm.PhoneNumber = input.PhoneNumber

	tx = repo.db.Save(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}

// SelectAll implements user.UserDataInterface
func (repo *userQuery) SelectAll() ([]user.Core, error) {
	var usersData []User
	tx := repo.db.Find(&usersData) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}
	// fmt.Println("users:", usersData)
	//mapping dari struct gorm model ke struct core (entity)
	var usersCore = ListModelToCore(usersData)

	return usersCore, nil
}

// SelectAll implements user.UserDataInterface

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(email string, password string) (dataLogin user.Core, err error) {
	var data User
	tx := repo.db.Where("email = ? and password = ?", email, password).Find(&data)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user.Core{}, errors.New("data not found")
	}
	dataLogin = ModelToCore(data)
	return dataLogin, nil
}

// SelectById implements user.UserDataInterface.
func (repo *userQuery) SelectById(id uint) (user.Core, error) {
	var result User
	tx := repo.db.First(&result, id)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user.Core{}, errors.New("data not found")
	}

	resultCore := ModelToCore(result)
	return resultCore, nil
}

// Insert implements user.UserDataInterface.
func (repo *userQuery) Insert(input user.Core) error {
	// mapping dari struct core to struct gorm model
	userGorm := CoreToModel(input)

	// simpan ke DB
	tx := repo.db.Create(&userGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
