package project

import (
	"github.com/Krab1o/taskboard/internal/converter/project"
	"github.com/Krab1o/taskboard/internal/converter/user"
	"github.com/Krab1o/taskboard/internal/model/dto"
	model "github.com/Krab1o/taskboard/internal/model/service"
)

func TaskModelToDTO(task *model.Task) *dto.Task {
	project := project.ProjectModelToDTO(task.Project)
	creator := user.UserModelToDTO(task.Creator)
	assignee := user.UserModelToDTO(task.Assigned)

	return &dto.Task{
		Id:           task.Id,
		Title:        task.Title,
		Description:  task.Description,
		Project:      project,
		CreatorUser:  creator,
		AssignedUser: assignee,
	}
}
