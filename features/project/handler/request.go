package handler

import (
	"royan/cleanarch/features/project"
)

type ProjectRequest struct {
	Name          string `json:"name" form:"name"`
	UserID        uint   `json:"user_id" form:"user_id"`
	DetailProject string `json:"detail_project" form:"detail_project"`
	Description   string `json:"description" form:"description"`
}

func RequestToCore(input ProjectRequest) project.Core {
	return project.Core{
		Name:          input.Name,
		DetailProject: input.DetailProject,
		Description:   input.Description,
	}
}
