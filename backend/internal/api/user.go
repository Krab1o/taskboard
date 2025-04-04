package api

import (
	"net/http"

	converter "github.com/Krab1o/taskboard/internal/converter/user"
	"github.com/Krab1o/taskboard/internal/model/dto"
	model "github.com/Krab1o/taskboard/internal/model/service"
	"github.com/gin-gonic/gin"
)

// @Tags			User
// @Summary		Get users list
// @Schemes		http
// @Description	Returns users in the system
// @Accept			json
// @Produce		json
// @Param		project		query		string	false	"Project Name"
// @Success		200		{array}	dto.User
// @Failure		400
// @Failure		500
// @Router			/users [get]
func (h *Handler) ListUsers(c *gin.Context) {
	ctx := c.Request.Context()
	projectTitle := c.Query("project")
	var users []model.User
	var err error
	if projectTitle == "" {
		users, err = h.userService.List(ctx)
	} else {
		users, err = h.userService.ListWithProject(ctx, projectTitle)
	}
	if err != nil {
		c.Status(http.StatusBadRequest)
	}

	dtoUsers := make([]dto.User, len(users))
	for i, user := range users {
		dtoUsers[i] = *converter.UserModelToDTO(&user)
	}
	c.JSON(http.StatusOK, dtoUsers)
}
