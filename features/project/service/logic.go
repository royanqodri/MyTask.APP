package service

import (
	"errors"
	"royan/cleanarch/features/project"
)

type projectService struct {
	projectData project.ProjectDataInterface
}

// Deletes implements project.ProjectServiceInterface
func (service *projectService) Deletes(id uint) error {
	err := service.projectData.Delete(id)
	return err
}

// Update implements project.ProjectServiceInterface
func (service *projectService) Update(id uint, input project.Core) error {
	err := service.projectData.Update(id, input)
	return err
}

// GetById implements project.ProjectServiceInterface
func (repo *projectService) GetById(id uint) (project.Core, error) {
	return repo.projectData.SelectById(id)
}

// GetAll implements project.ProjectServiceInterface
func (service *projectService) GetAll() ([]project.Core, error) {
	result, err := service.projectData.SelectAll()
	return result, err
}

// Create implements project.ProjectServiceInterface
func (service *projectService) Create(input project.Core) error {
	if input.Name == "" || input.DetailProject == "" || input.Description == "" {
		return errors.New("validation error. name/detail/deskiption required")
	}
	err := service.projectData.Insert(input)
	return err
}

func New(repo project.ProjectDataInterface) project.ProjectServiceInterface {
	return &projectService{
		projectData: repo,
	}
}
