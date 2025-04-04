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
	ErrListTasksRepoMessage = errors.New("failed to list tasks repository")
)

const (
	UserTableAliasCreator  = "creator"
	UserTableAliasAssigned = "assigned"
)

func (r *TaskRepository) ListTasksByProjectId(
	ctx context.Context,
	projectTitle string,
) ([]model.Task, error) {
	query, args, err := squirrel.Select(
		col(TaskTableName, TaskColumnId),

		col(TaskTableName, TaskColumnIdProject),
		col(ProjectTableName, ProjectColumnTitle),

		col(TaskTableName, TaskColumnIdCreator),
		fmt.Sprintf("%s AS creator_given_name", col(UserTableAliasCreator, UserColumnGivenName)),
		fmt.Sprintf("%s AS creator_surname", col(UserTableAliasCreator, UserColumnSurname)),

		col(TaskTableName, TaskColumnIdAssigned),
		fmt.Sprintf("%s AS assigned_given_name", col(UserTableAliasAssigned, UserColumnGivenName)),
		fmt.Sprintf("%s AS assigned_surname", col(UserTableAliasAssigned, UserColumnSurname)),

		col(TaskTableName, TaskColumnIdPriority),
		col(TaskTableName, TaskColumnIdStatus),
		col(TaskTableName, TaskColumnTitle),
		col(TaskTableName, TaskColumnDescription),
	).
		PlaceholderFormat(squirrel.Dollar).
		From(TaskTableName).
		LeftJoin(fmt.Sprintf("%s ON %s = %s",
			ProjectTableName,
			col(TaskTableName, TaskColumnIdProject),
			col(ProjectTableName, ProjectColumnId),
		)).
		LeftJoin(fmt.Sprintf("%s ON %s = %s",
			StatusTableName,
			col(TaskTableName, TaskColumnIdStatus),
			col(StatusTableName, StatusColumnId),
		)).
		LeftJoin(fmt.Sprintf("%s ON %s = %s",
			PriorityTableName,
			col(TaskTableName, TaskColumnIdPriority),
			col(PriorityTableName, PriorityColumnId),
		)).
		LeftJoin(fmt.Sprintf("%s AS %s ON %s = %s",
			UserTableName,
			UserTableAliasCreator,
			col(TaskTableName, TaskColumnIdCreator),
			col(UserTableAliasCreator, UserColumnId),
		)).
		LeftJoin(fmt.Sprintf("%s AS %s ON %s = %s",
			UserTableName,
			UserTableAliasAssigned,
			col(TaskTableName, TaskColumnIdAssigned),
			col(UserTableAliasAssigned, UserColumnId),
		)).
		Where(squirrel.Eq{col(ProjectTableName, ProjectColumnTitle): projectTitle}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUserListMessage, err)
	}

	log.Println(query)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUserListMessage, err)
	}

	tasks := []model.Task{}
	task := &model.Task{}
	for rows.Next() {
		task = &model.Task{
			Project:  &model.Project{},
			Creator:  &model.User{},
			Assigned: &model.User{},
		}
		err = rows.Scan(
			&task.Id,
			&task.Project.Id,
			&task.Project.Title,

			&task.Creator.Id,
			&task.Creator.GivenName,
			&task.Creator.Surname,

			&task.Assigned.Id,
			&task.Assigned.GivenName,
			&task.Assigned.Surname,

			&task.PriorityId,
			&task.StatusId,
			&task.Title,
			&task.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrUserListMessage, err)
		}

		tasks = append(tasks, *task)
	}
	return tasks, nil
}
