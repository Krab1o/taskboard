package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	model "github.com/Krab1o/taskboard/internal/model/service"
	"github.com/Masterminds/squirrel"
)

var (
	ErrUserListMessage = errors.New("failed to list users repository")
)

func (r *UserRepository) List(ctx context.Context) ([]model.User, error) {
	query, args, err := squirrel.Select(
		col(UserTableName, UserColumnGivenName),
		col(UserTableName, UserColumnSurname),
	).
		PlaceholderFormat(squirrel.Dollar).
		From(
			UserTableName,
		).ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUserListMessage, err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUserListMessage, err)
	}

	users := []model.User{}
	user := &model.User{}
	for rows.Next() {
		user = &model.User{}
		err = rows.Scan(
			&user.GivenName,
			&user.Surname,
		)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrUserListMessage, err)
		}

		users = append(users, *user)
	}

	return users, nil
}

func (r *UserRepository) ListWithProject(
	ctx context.Context,
	projectTitle string,
) ([]model.User, error) {
	query, args, err := squirrel.Select(
		col(UserTableName, UserColumnGivenName),
		col(UserTableName, UserColumnSurname),
	).
		PlaceholderFormat(squirrel.Dollar).
		From(
			UserTableName,
		).LeftJoin(fmt.Sprintf("%s ON %s = %s",
		UserProjectTableName,
		col(UserProjectTableName, UserProjectColumnUserId),
		col(UserTableName, UserColumnId),
	)).LeftJoin(fmt.Sprintf("%s ON %s = %s",
		ProjectTableName,
		col(ProjectTableName, ProjectColumnId),
		col(UserProjectTableName, UserProjectColumnProjectId),
	)).Where(
		squirrel.Eq{col(ProjectTableName, ProjectColumnTitle): projectTitle},
	).ToSql()

	log.Println(query)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUserListMessage, err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUserListMessage, err)
	}

	users := []model.User{}
	user := &model.User{}
	for rows.Next() {
		user = &model.User{}
		err = rows.Scan(
			&user.GivenName,
			&user.Surname,
		)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrUserListMessage, err)
		}

		users = append(users, *user)
	}

	return users, nil
}
