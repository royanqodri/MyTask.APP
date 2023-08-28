package data

import (
	"errors"
	"royan/cleanarch/features/project"

	"gorm.io/gorm"
)

type projectQuery struct {
	db *gorm.DB
}

// Delete implements project.ProjectDataInterface
func (repo *projectQuery) Delete(id uint) error {
	var projectGorm Project
	tx := repo.db.First(&projectGorm, id)
	if tx.Error != nil {
		return tx.Error
	}

	// Hapus pengguna dari database
	tx = repo.db.Delete(&projectGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("data not found to deleted")

	}
	return nil
}

// Update implements project.ProjectDataInterface
func (repo *projectQuery) Update(id uint, input project.Core) error {
	var projectGorm Project
	tx := repo.db.First(&projectGorm, id)
	if tx.Error != nil {
		return tx.Error
	}

	// Perbarui properti pengguna dengan data dari input
	projectGorm.Name = input.Name
	projectGorm.UserID = input.UserID
	projectGorm.DetailProject = input.DetailProject
	projectGorm.Description = input.Description

	tx = repo.db.Save(&projectGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}

// SelectById implements project.ProjectDataInterface
func (repo *projectQuery) SelectById(id uint) (project.Core, error) {
	var result Project
	tx := repo.db.First(&result, id)
	if tx.Error != nil {
		return project.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return project.Core{}, errors.New("data not found")
	}

	resultCore := ModelToCore(result)
	return resultCore, nil
}

// SelectAll implements project.ProjectDataInterface
func (repo *projectQuery) SelectAll() ([]project.Core, error) {
	var projectData []Project
	tx := repo.db.Find(&projectData) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}
	// fmt.Println("users:", usersData)
	//mapping dari struct gorm model ke struct core (entity)
	var projectCore = ListModelToCore(projectData)

	return projectCore, nil
}

// Insert implements project.ProjectDataInterface
func (repo *projectQuery) Insert(input project.Core) error {
	projectGorm := CoreToModel(input)

	// simpan ke DB
	tx := repo.db.Create(&projectGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func New(db *gorm.DB) project.ProjectDataInterface {
	return &projectQuery{
		db: db,
	}
}
