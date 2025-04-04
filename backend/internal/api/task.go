package api

import (
	"net/http"

	converter "github.com/Krab1o/taskboard/internal/converter/task"
	"github.com/Krab1o/taskboard/internal/model/dto"
	model "github.com/Krab1o/taskboard/internal/model/service"
	"github.com/gin-gonic/gin"
)

// @Tags			Tasks
// @Summary		Get tasks list
// @Schemes		http
// @Description	Returns tasks in the system
// @Accept			json
// @Produce		json
// @Param		project		query		string	true	"Project Name"
// @Success		200		{array}	dto.Task
// @Failure		400
// @Failure		500
// @Router			/tasks [get]
func (h *Handler) ListTasks(c *gin.Context) {
	ctx := c.Request.Context()
	projectTitle := c.Query("project")
	var tasks []model.Task
	var err error

	tasks, err = h.taskService.ListWithProject(ctx, projectTitle)

	if err != nil {
		c.Status(http.StatusBadRequest)
	}

	dtoTasks := make([]dto.Task, len(tasks))
	for i, task := range tasks {
		dtoTasks[i] = *converter.TaskModelToDTO(&task)
	}
	c.JSON(http.StatusOK, dtoTasks)
}

func (h *Handler) CreateTask(c *gin.Context) {

}
