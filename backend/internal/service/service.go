package service

import (
	"github.com/Krab1o/taskboard/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

type ProjectService struct {
	projectRepository *repository.ProjectRepository
}

type TaskService struct {
	taskRepository *repository.TaskRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func NewProjectService(projectRepository *repository.ProjectRepository) *ProjectService {
	return &ProjectService{
		projectRepository: projectRepository,
	}
}

func NewTaskService(taskRepository *repository.TaskRepository) *TaskService {
	return &TaskService{
		taskRepository: taskRepository,
	}
}
