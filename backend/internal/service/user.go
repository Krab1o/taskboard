package service

import (
	"context"
	"errors"
	"fmt"

	model "github.com/Krab1o/taskboard/internal/model/service"
)

var (
	ErrMessage = errors.New("failed to list users service")
)

func (s *UserService) List(ctx context.Context) ([]model.User, error) {
	users, err := s.userRepository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrMessage, err)
	}
	return users, nil
}

func (s *UserService) ListWithProject(
	ctx context.Context,
	projectTitle string,
) ([]model.User, error) {
	users, err := s.userRepository.ListWithProject(ctx, projectTitle)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrMessage, err)
	}
	return users, nil
}
