package model

type Task struct {
	Id          uint64
	StatusId    Status
	PriorityId  Priority
	Title       string
	Description string
	Project     *Project
	Creator     *User
	Assigned    *User
}
