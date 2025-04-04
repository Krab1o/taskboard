package dto

type NewTask struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	ProjectId    uint64 `json:"projectId"`
	CreatorUser  uint64 `json:"creatorId"`
	AssignedUser uint64 `json:"assignedId"`
}

type Task struct {
	Id           uint64
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Project      *Project `json:"project"`
	CreatorUser  *User    `json:"creatorUser"`
	AssignedUser *User    `json:"assignedUser"`
}
