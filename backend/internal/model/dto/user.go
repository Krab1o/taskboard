package dto

type User struct {
	Id        uint64 `json:"id"`
	GivenName string `json:"givenName"`
	Surname   string `json:"surname"`
}
