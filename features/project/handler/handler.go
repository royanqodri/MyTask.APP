package handler

import (
	"net/http"
	"royan/cleanarch/app/middlewares"
	"royan/cleanarch/features/project"
	"royan/cleanarch/helpers"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	projectService project.ProjectServiceInterface
}

func NewProjectHandler(service project.ProjectServiceInterface) *ProjectHandler {
	return &ProjectHandler{
		projectService: service,
	}
}

func (handler *ProjectHandler) CreateProject(c *gin.Context) {
	// Ambil ID pengguna dari token
	id := middlewares.ExtractTokenUserId(c)

	var projectInput ProjectRequest
	if err := c.BindJSON(&projectInput); err != nil {
		helpers.WebResponse(c, http.StatusBadRequest, "Error binding data", nil)
		return
	}

	// Mapping dari struct request ke struct core
	projectCore := RequestToCore(projectInput)
	projectCore.UserID = uint(id)

	err := handler.projectService.Create(projectCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			helpers.WebResponse(c, http.StatusBadRequest, err.Error(), nil)
		} else {
			helpers.WebResponse(c, http.StatusInternalServerError, "Error inserting data", nil)
		}
		return
	}

	helpers.WebResponse(c, http.StatusCreated, "Success inserting data", nil)
}

func (handler *ProjectHandler) GetAllProjects(c *gin.Context) {
	result, err := handler.projectService.GetAll()
	if err != nil {
		helpers.WebResponse(c, http.StatusInternalServerError, "Error reading data", nil)
		return
	}

	// Mapping dari struct core ke struct response
	var projectResponse []ProjectResponse
	for _, value := range result {
		projectResponse = append(projectResponse, ProjectResponse{
			ID:            value.ID,
			Name:          value.Name,
			UserID:        value.UserID,
			DetailProject: value.DetailProject,
			Description:   value.Description,
			CreatedAt:     value.CreatedAt,
		})
	}

	helpers.WebResponse(c, http.StatusOK, "Success reading data", projectResponse)
}

func (handler *ProjectHandler) GetProjectById(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("projectid"))
	if err != nil || projectID <= 0 {
		helpers.WebResponse(c, http.StatusBadRequest, "Invalid project ID", nil)
		return
	}

	result, err := handler.projectService.GetById(uint(projectID))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			helpers.WebResponse(c, http.StatusBadRequest, err.Error(), nil)
		} else {
			helpers.WebResponse(c, http.StatusInternalServerError, "Error reading data", nil)
		}
		return
	}

	projectResponse := ProjectResponse{
		ID:            result.ID,
		Name:          result.Name,
		UserID:        result.UserID,
		DetailProject: result.DetailProject,
		Description:   result.Description,
		CreatedAt:     result.CreatedAt,
	}

	helpers.WebResponse(c, http.StatusOK, "Success reading data", projectResponse)
}

func (handler *ProjectHandler) UpdateProject(c *gin.Context) {
	// Ambil ID pengguna dari token
	id := middlewares.ExtractTokenUserId(c)

	projectID, err := strconv.Atoi(c.Param("projectid"))
	if err != nil || projectID <= 0 {
		helpers.WebResponse(c, http.StatusBadRequest, "Invalid project ID", nil)
		return
	}

	var projectInput ProjectRequest
	if err := c.BindJSON(&projectInput); err != nil {
		helpers.WebResponse(c, http.StatusBadRequest, "Error binding data", nil)
		return
	}

	// Mapping dari struct request ke struct core
	projectCore := RequestToCore(projectInput)
	projectCore.UserID = uint(id)

	err = handler.projectService.Update(uint(projectID), projectCore)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			helpers.WebResponse(c, http.StatusNotFound, "Project not found", nil)
		} else {
			helpers.WebResponse(c, http.StatusInternalServerError, "Error updating project", nil)
		}
		return
	}

	helpers.WebResponse(c, http.StatusOK, "Project updated successfully", nil)
}

func (handler *ProjectHandler) DeleteProject(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("projectid"))
	if err != nil || projectID <= 0 {
		helpers.WebResponse(c, http.StatusBadRequest, "Invalid project ID", nil)
		return
	}

	err = handler.projectService.Deletes(uint(projectID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			helpers.WebResponse(c, http.StatusNotFound, "Project not found", nil)
		} else {
			helpers.WebResponse(c, http.StatusInternalServerError, "Error deleting project", nil)
		}
		return
	}

	helpers.WebResponse(c, http.StatusOK, "Project deleted successfully", nil)
}
