package handler

import (
	"net/http"
	"royan/cleanarch/app/middlewares"
	"royan/cleanarch/features/task"
	"royan/cleanarch/helpers"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskService task.TaskServiceInterface
}

func New(service task.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{
		taskService: service, // Mengganti projectService dengan service
	}
}

func (handler *TaskHandler) CreateTask(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	taskInput := new(TaskRequest)
	errBind := c.Bind(&taskInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	//mapping dari struct request to struct core
	taskCore := RequestToCore(*taskInput)
	taskCore.ProjectID = uint(id)
	err := handler.taskService.Create(taskCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))

		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Item created successfully", nil))
}

func (handler *TaskHandler) DeleteTask(c echo.Context) error {
	// Ambil `user_id` dari parameter URL
	taskID, err := strconv.Atoi(c.Param("taskid"))
	if err != nil || taskID <= 0 {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid user_id", nil))
	}

	// Panggil fungsi service untuk menghapus pengguna berdasarkan `user_id`
	err = handler.taskService.Deletes(uint(taskID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "taks not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error deleting task", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Task deleted successfully", nil))
}

func (handler *TaskHandler) UpdateTask(c echo.Context) error {
	// Ambil `user_id` dari parameter URL
	id := middlewares.ExtractTokenUserId(c)
	taskID, err := strconv.Atoi(c.Param("taskid"))
	if err != nil || taskID <= 0 {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid user_id", nil))
	}

	// Ambil data pengguna yang akan diperbarui dari permintaan JSON
	taskInput := new(TaskRequest)
	errBind := c.Bind(&taskInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Error binding data", nil))
	}

	// Mapping data dari UserRequest ke struct Core
	taskCore := RequestToCore(*taskInput)
	taskCore.ProjectID = uint(id)

	// Panggil fungsi service untuk memperbarui pengguna
	err = handler.taskService.Update(uint(taskID), taskCore)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "Task not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error updating task", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Task updated successfully", nil))
}
