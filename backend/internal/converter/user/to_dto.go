package user

import (
	"github.com/Krab1o/taskboard/internal/model/dto"
	model "github.com/Krab1o/taskboard/internal/model/service"
)

func UserModelToDTO(user *model.User) *dto.User {
	return &dto.User{
		Id:        user.Id,
		GivenName: user.GivenName,
		Surname:   user.Surname,
	}
}
