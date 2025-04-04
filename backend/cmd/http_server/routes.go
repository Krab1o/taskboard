package main

import (
	docs "github.com/Krab1o/taskboard/docs"
	"github.com/Krab1o/taskboard/internal/api"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Taskboard
//	@version		0.1
//	@description	API for taskboard application
//	@contact.name	Grigory

//	@host		localhost:8080
//	@BasePath	/api
//	@accept		json
//	@produce	json
//	@schemes	http

//	@tag.name			User
//	@tag.description	User endpoints

//	@tag.name			Projects
//	@tag.description	Project endpoints

// @tag.name			Tasks
// @tag.description	Tasks endpoints
func setupRoutes(s *gin.Engine, handler *api.Handler) {
	docs.SwaggerInfo.BasePath = "/api"

	s.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	apiGroup := s.Group("/api")
	{
		userGroup := apiGroup.Group("/users")
		{
			userGroup.GET("", handler.ListUsers)
		}
		projectGroup := apiGroup.Group("/projects")
		{
			projectGroup.GET("", handler.ListProjects)
		}
		taskGroup := apiGroup.Group("/tasks")
		{
			taskGroup.GET("", handler.ListTasks)
			taskGroup.POST("", handler.CreateTask)
		}
	}
}
