package api

import "github.com/Krab1o/taskboard/internal/service"

type Handler struct {
	userService    *service.UserService
	projectService *service.ProjectService
	taskService    *service.TaskService
}

func NewHandler(
	userService *service.UserService,
	projectService *service.ProjectService,
	taskService *service.TaskService,
) *Handler {
	return &Handler{
		userService:    userService,
		projectService: projectService,
		taskService:    taskService,
	}
}
