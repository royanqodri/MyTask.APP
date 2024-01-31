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

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *gin.Engine) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	projectData := _projectData.New(db)
	projectService := _projectService.New(projectData)
	projectHandlerAPI := _projectHandler.NewProjectHandler(projectService)

	taskData := _taskData.New(db)
	taskService := _taskService.New(taskData)
	taskHandlerAPI := _taskHandler.New(taskService)

	api := e.Group("/api") // Grupkan semua rute API di bawah "/api"

	api.POST("/login", userHandlerAPI.Login)

	api.GET("/users", userHandlerAPI.GetAllUser, middlewares.JWTMiddleware())
	api.GET("/users/:user_id", userHandlerAPI.GetUserById)
	api.POST("/users", userHandlerAPI.CreateUser, middlewares.JWTMiddleware())
	api.PUT("/users", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())
	api.DELETE("/users", userHandlerAPI.DeleteUser, middlewares.JWTMiddleware())

	api.POST("/projects", projectHandlerAPI.CreateProject, middlewares.JWTMiddleware())
	api.GET("/projects", projectHandlerAPI.GetAllProjects, middlewares.JWTMiddleware())
	api.GET("/projects/:projectid", projectHandlerAPI.GetProjectById, middlewares.JWTMiddleware())
	api.PUT("/projects/:projectid", projectHandlerAPI.UpdateProject, middlewares.JWTMiddleware())
	api.DELETE("/projects/:projectid", projectHandlerAPI.DeleteProject, middlewares.JWTMiddleware())

	api.POST("/tasks", taskHandlerAPI.CreateTask, middlewares.JWTMiddleware())
	api.DELETE("/tasks/:taskid", taskHandlerAPI.DeleteTask, middlewares.JWTMiddleware())
	api.PUT("/tasks/:taskid", taskHandlerAPI.UpdateTask, middlewares.JWTMiddleware())

	// Rute non-API
	e.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
}
