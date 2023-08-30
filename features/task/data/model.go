package data

import (
	_projectData "royan/cleanarch/features/project/data"
	"royan/cleanarch/features/task"

	"gorm.io/gorm"
)

// struct item gorm model
type Task struct {
	gorm.Model
	// ID          uint `gorm:"primaryKey"`
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
	// DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name       string
	ProjectID  uint
	DetailTask string
	Status     string
	Project    _projectData.Project `gorm:"foreignKey:ProjectID"`
}

func CoreToModel(dataCore task.Core) Task {
	return Task{
		Name:       dataCore.Name,
		ProjectID:  dataCore.ProjectID,
		DetailTask: dataCore.DetailTask,
		Status:     dataCore.Status,
		Project:    _projectData.Project{},
	}
}

// mapping struct model to struct core
func ModelToCore(dataModel Task) task.Core {
	return task.Core{
		ID:         dataModel.ID,
		Name:       dataModel.Name,
		ProjectID:  dataModel.ProjectID,
		DetailTask: dataModel.DetailTask,
		Status:     dataModel.Status,
		CreatedAt:  dataModel.CreatedAt,
		UpdatedAt:  dataModel.UpdatedAt,
		Project:    task.ProjectCore{},
	}
}
