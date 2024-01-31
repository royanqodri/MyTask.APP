package handler

import (
	"net/http"
	"royan/cleanarch/app/middlewares"
	"royan/cleanarch/features/task"
	"royan/cleanarch/helpers"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService task.TaskServiceInterface
}

func New(service task.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{
		taskService: service,
	}
}

func (handler *TaskHandler) CreateTask(c *gin.Context) {
	id := middlewares.ExtractTokenUserId(c)
	taskInput := new(TaskRequest)
	errBind := c.Bind(&taskInput)
	if errBind != nil {
		helpers.WebResponse(c, http.StatusBadRequest, "error bind data. data not valid", nil)
		return
	}

	taskCore := RequestToCore(*taskInput)
	taskCore.ProjectID = uint(id)

	err := handler.taskService.Create(taskCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			helpers.WebResponse(c, http.StatusBadRequest, err.Error(), nil)
		} else {
			helpers.WebResponse(c, http.StatusInternalServerError, "error insert data", nil)
		}
		return
	}

	helpers.WebResponse(c, http.StatusOK, "Item created successfully", nil)
}

func (handler *TaskHandler) DeleteTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("taskid"))
	if err != nil || taskID <= 0 {
		helpers.WebResponse(c, http.StatusBadRequest, "Invalid taskid", nil)
		return
	}

	err = handler.taskService.Deletes(uint(taskID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			helpers.WebResponse(c, http.StatusNotFound, "Task not found", nil)
		} else {
			helpers.WebResponse(c, http.StatusInternalServerError, "Error deleting task", nil)
		}
		return
	}

	helpers.WebResponse(c, http.StatusOK, "Task deleted successfully", nil)
}

func (handler *TaskHandler) UpdateTask(c *gin.Context) {
	id := middlewares.ExtractTokenUserId(c)
	taskID, err := strconv.Atoi(c.Param("taskid"))
	if err != nil || taskID <= 0 {
		helpers.WebResponse(c, http.StatusBadRequest, "Invalid taskid", nil)
		return
	}

	taskInput := new(TaskRequest)
	errBind := c.Bind(&taskInput)
	if errBind != nil {
		helpers.WebResponse(c, http.StatusBadRequest, "Error binding data", nil)
		return
	}

	taskCore := RequestToCore(*taskInput)
	taskCore.ProjectID = uint(id)

	err = handler.taskService.Update(uint(taskID), taskCore)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			helpers.WebResponse(c, http.StatusNotFound, "Task not found", nil)
		} else {
			helpers.WebResponse(c, http.StatusInternalServerError, "Error updating task", nil)
		}
		return
	}

	helpers.WebResponse(c, http.StatusOK, "Task updated successfully", nil)
}
