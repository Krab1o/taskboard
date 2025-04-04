package service

import (
	"context"
	"errors"
	"fmt"

	model "github.com/Krab1o/taskboard/internal/model/service"
)

var (
	ErrListProjectsServMessage = errors.New("failed to list projects service")
)

func (s *ProjectService) ListProjects(ctx context.Context) ([]model.Project, error) {
	projects, err := s.projectRepository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrListTasksServMessage, err)
	}
	return projects, nil
}
