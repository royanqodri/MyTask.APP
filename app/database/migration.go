package database

import (
	_projectData "royan/cleanarch/features/project/data"
	_taskData "royan/cleanarch/features/task/data"
	_userData "royan/cleanarch/features/user/data"

	"gorm.io/gorm"
)

// db migration
func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_userData.User{})
	db.AutoMigrate(&_projectData.Project{})
	db.AutoMigrate(&_taskData.Task{})

	/*
		TODO 2:
		migrate struct item
	*/
}
