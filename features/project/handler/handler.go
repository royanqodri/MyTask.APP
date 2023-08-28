package handler

import (
	"net/http"
	"royan/cleanarch/app/middlewares"
	"royan/cleanarch/features/project"
	"royan/cleanarch/helpers"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	projectService project.ProjectServiceInterface
}

func New(service project.ProjectServiceInterface) *ProjectHandler {
	return &ProjectHandler{
		projectService: service, // Mengganti projectService dengan service
	}
}

func (handler *ProjectHandler) CreateProject(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	userInput := new(ProjectRequest)
	errBind := c.Bind(&userInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	//mapping dari struct request to struct core
	projectCore := RequestToCore(*userInput)
	projectCore.UserID = uint(id)
	err := handler.projectService.Create(projectCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))

		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success insert data", nil))
}

func (handler *ProjectHandler) GetAllProject(c echo.Context) error {
	result, err := handler.projectService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	// mapping dari struct core to struct response
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
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", projectResponse))

}

func (handler *ProjectHandler) GetProjectById(c echo.Context) error {
	// Ambil ID proyek dari parameter URL
	id := c.Param("projectid")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil || idConv <= 0 {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid project ID", nil))
	}

	// Panggil fungsi service untuk mendapatkan detail proyek
	result, err := handler.projectService.GetById(uint(idConv))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error reading data", nil))
		}
	}

	resultResponse := ProjectResponse{
		ID:            result.ID,
		Name:          result.Name,
		UserID:        result.UserID,
		DetailProject: result.DetailProject,
		Description:   result.Description,
		CreatedAt:     result.CreatedAt,
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success reading data", resultResponse))
}

func (handler *ProjectHandler) UpdateProject(c echo.Context) error {
	// Ambil `user_id` dari parameter URL
	id := middlewares.ExtractTokenUserId(c)
	userID, err := strconv.Atoi(c.Param("projectid"))
	if err != nil || userID <= 0 {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid user_id", nil))
	}

	// Ambil data pengguna yang akan diperbarui dari permintaan JSON
	projectInput := new(ProjectRequest)
	errBind := c.Bind(&projectInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Error binding data", nil))
	}

	// Mapping data dari UserRequest ke struct Core
	projectCore := RequestToCore(*projectInput)
	projectCore.UserID = uint(id)

	// Panggil fungsi service untuk memperbarui pengguna
	err = handler.projectService.Update(uint(userID), projectCore)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "project not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error updating project", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Project updated successfully", nil))
}

func (handler *ProjectHandler) DeleteProject(c echo.Context) error {
	// Ambil `user_id` dari parameter URL
	projectID, err := strconv.Atoi(c.Param("projectid"))
	if err != nil || projectID <= 0 {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid user_id", nil))
	}

	// Panggil fungsi service untuk menghapus pengguna berdasarkan `user_id`
	err = handler.projectService.Deletes(uint(projectID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "Project not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error deleting Project", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Project deleted successfully", nil))
}
