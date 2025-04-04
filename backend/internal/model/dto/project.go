package dto

type Project struct {
	Id        uint64 `json:"id"`
	Title     string `json:"title"`
	TaskCount uint64 `json:"taskCount"`
}
