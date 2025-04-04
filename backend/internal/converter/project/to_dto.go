package project

import (
	"github.com/Krab1o/taskboard/internal/model/dto"
	model "github.com/Krab1o/taskboard/internal/model/service"
)

func ProjectModelToDTO(project *model.Project) *dto.Project {
	return &dto.Project{
		Id:        project.Id,
		Title:     project.Title,
		TaskCount: project.TaskCount,
	}
}
