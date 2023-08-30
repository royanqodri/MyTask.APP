package data

import (
	"errors"
	"royan/cleanarch/features/task"

	"gorm.io/gorm"
)

type taskQuery struct {
	db *gorm.DB
}

// Update implements task.TaskDataInterface
func (repo *taskQuery) Update(id uint, input task.Core) error {
	var taskGorm Task
	tx := repo.db.First(&taskGorm, id)
	if tx.Error != nil {
		return tx.Error
	}

	// Perbarui properti pengguna dengan data dari input
	taskGorm.Name = input.Name
	taskGorm.ProjectID = input.ProjectID
	taskGorm.DetailTask = input.DetailTask
	taskGorm.Status = input.Status

	tx = repo.db.Save(&taskGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}

// Delete implements task.TaskDataInterface
func (repo *taskQuery) Delete(id uint) error {
	var taskGorm Task
	tx := repo.db.First(&taskGorm, id)
	if tx.Error != nil {
		return tx.Error
	}

	// Hapus pengguna dari database
	tx = repo.db.Delete(&taskGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("data not found to deleted")

	}
	return nil
}

// Insert implements task.TaskDataInterface
func (repo *taskQuery) Insert(input task.Core) error {

	taskGorm := CoreToModel(input)

	// simpan ke DB
	tx := repo.db.Create(&taskGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func New(db *gorm.DB) task.TaskDataInterface {
	return &taskQuery{
		db: db,
	}
}
