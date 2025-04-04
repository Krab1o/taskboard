package repository

import (
	"context"
	"fmt"

	model "github.com/Krab1o/taskboard/internal/model/service"
	"github.com/Masterminds/squirrel"
)

func (r *ProjectRepository) List(ctx context.Context) ([]model.Project, error) {
	query, args, err := squirrel.Select(
		col(ProjectTableName, ProjectColumnId),
		col(ProjectTableName, ProjectColumnTitle),
		fmt.Sprintf("COUNT(%s)", col(TaskTableName, TaskColumnId)),
	).
		PlaceholderFormat(squirrel.Dollar).
		From(ProjectTableName).
		LeftJoin(fmt.Sprintf("%s ON %s = %s",
			TaskTableName,
			col(TaskTableName, TaskColumnIdProject),
			col(ProjectTableName, ProjectColumnId),
		)).
		GroupBy(
			col(ProjectTableName, ProjectColumnId),
			col(ProjectTableName, ProjectColumnTitle)).
		OrderBy(col(ProjectTableName, ProjectColumnTitle)).
		ToSql()

	squirrel.Select().
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

	projects := []model.Project{}
	project := &model.Project{}
	for rows.Next() {
		project = &model.Project{}
		err = rows.Scan(
			&project.Id,
			&project.Title,
			&project.TaskCount,
		)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrUserListMessage, err)
		}

		projects = append(projects, *project)
	}

	return projects, nil
}
