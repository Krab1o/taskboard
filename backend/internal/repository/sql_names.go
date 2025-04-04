package repository

import "fmt"

func col(table, column string) string {
	return fmt.Sprintf("%s.%s", table, column)
}

const (
	UserTableName       = "user_"
	UserColumnId        = "id_"
	UserColumnGivenName = "given_name_"
	UserColumnSurname   = "surname_"

	TaskTableName         = "task_"
	TaskColumnId          = "id_"
	TaskColumnIdProject   = "id_project_"
	TaskColumnIdStatus    = "id_status_"
	TaskColumnIdPriority  = "id_priority_"
	TaskColumnTitle       = "title_"
	TaskColumnDescription = "description_"
	TaskColumnIdCreator   = "id_creator_"
	TaskColumnIdAssigned  = "id_assigned_"

	ProjectTableName   = "project_"
	ProjectColumnId    = "id_"
	ProjectColumnTitle = "title_"

	UserProjectTableName       = "user_project_"
	UserProjectColumnUserId    = "id_user_"
	UserProjectColumnProjectId = "id_project_"

	StatusTableName   = "status_"
	StatusColumnId    = "id_"
	StatusColumnTitle = "title_"

	PriorityTableName   = "priority_"
	PriorityColumnId    = "id_"
	PriorityColumnTitle = "title_"
)
