package router

import (
	"royan/cleanarch/app/middlewares"
	_userData "royan/cleanarch/features/user/data"
	_userHandler "royan/cleanarch/features/user/handler"
	_userService "royan/cleanarch/features/user/service"

	_projectData "royan/cleanarch/features/project/data"
	_projectHandler "royan/cleanarch/features/project/handler"
	_projectService "royan/cleanarch/features/project/service"

	_taskData "royan/cleanarch/features/task/data"
	_taskHandler "royan/cleanarch/features/task/handler"
	_taskService "royan/cleanarch/features/task/service"

	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	projectData := _projectData.New(db)
	projectService := _projectService.New(projectData)
	projectHandlerAPI := _projectHandler.New(projectService)

	taskData := _taskData.New(db)
	taskService := _taskService.New(taskData)
	taskHandlerAPI := _taskHandler.New(taskService)

	e.POST("/login", userHandlerAPI.Login)

	e.GET("/users", userHandlerAPI.GetAllUser, middlewares.JWTMiddleware())
	e.GET("/users/:user_id", userHandlerAPI.GetUserById)
	e.POST("/users", userHandlerAPI.CreateUser, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.DeleteUser, middlewares.JWTMiddleware())

	e.POST("/projects", projectHandlerAPI.CreateProject, middlewares.JWTMiddleware())
	e.GET("/projects", projectHandlerAPI.GetAllProject, middlewares.JWTMiddleware())
	e.GET("/projects/:projectid", projectHandlerAPI.GetProjectById, middlewares.JWTMiddleware())
	e.PUT("/projects/:projectid", projectHandlerAPI.UpdateProject, middlewares.JWTMiddleware())
	e.DELETE("/projects/:projectid", projectHandlerAPI.DeleteProject, middlewares.JWTMiddleware())

	e.POST("/tasks", taskHandlerAPI.CreateTask, middlewares.JWTMiddleware())
	e.DELETE("/tasks/:taskid", taskHandlerAPI.DeleteTask, middlewares.JWTMiddleware())
	e.PUT("/tasks/:taskid", taskHandlerAPI.UpdateTask, middlewares.JWTMiddleware())

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "hello world",
		})
	})
}
