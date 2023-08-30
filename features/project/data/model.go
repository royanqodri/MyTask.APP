package data

import (
	"royan/cleanarch/features/project"
	_userData "royan/cleanarch/features/user/data"

	"gorm.io/gorm"
)

// struct item gorm model
type Project struct {
	gorm.Model
	// ID          uint `gorm:"primaryKey"`
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
	// DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name          string
	UserID        uint
	DetailProject string
	Description   string
	User          _userData.User `gorm:"foreignKey:UserID"`
}

func CoreToModel(dataCore project.Core) Project {
	return Project{
		Name:          dataCore.Name,
		UserID:        dataCore.UserID,
		DetailProject: dataCore.DetailProject,
		Description:   dataCore.Description,
		User:          _userData.User{},
	}
}

// mapping struct model to struct core
func ModelToCore(dataModel Project) project.Core {
	return project.Core{
		ID:            dataModel.ID,
		Name:          dataModel.Name,
		UserID:        dataModel.UserID,
		DetailProject: dataModel.DetailProject,
		Description:   dataModel.Description,
		CreatedAt:     dataModel.CreatedAt,
		UpdatedAt:     dataModel.UpdatedAt,
		User:          project.UserCore{},
	}
}

func ListModelToCore(dataModel []Project) []project.Core {
	var result []project.Core
	for _, v := range dataModel {
		result = append(result, ModelToCore(v))
	}
	return result
}
