package service

import (
	"context"
	"errors"
	"fmt"

	model "github.com/Krab1o/taskboard/internal/model/service"
)

var (
	ErrListTasksServMessage = errors.New("failed to list tasks service")
)

func (s *TaskService) ListWithProject(
	ctx context.Context,
	projectTitle string,
) ([]model.Task, error) {
	tasks, err := s.taskRepository.ListTasksByProjectId(ctx, projectTitle)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrListTasksServMessage, err)
	}
	return tasks, nil
}
