package api

import (
	"net/http"

	converter "github.com/Krab1o/taskboard/internal/converter/project"
	"github.com/Krab1o/taskboard/internal/model/dto"
	"github.com/gin-gonic/gin"
)

// @Tags			Project
// @Summary		Get projects list
// @Schemes		http
// @Description	Returns projects in the system
// @Accept			json
// @Produce		json
// @Success		200		{array}	dto.Project
// @Failure		400
// @Failure		500
// @Router			/projects [get]
func (h *Handler) ListProjects(c *gin.Context) {
	ctx := c.Request.Context()
	projects, err := h.projectService.ListProjects(ctx)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}

	dtoProjects := make([]dto.Project, len(projects))
	for i, project := range projects {
		dtoProjects[i] = *converter.ProjectModelToDTO(&project)
	}
	c.JSON(http.StatusOK, dtoProjects)
}

func (h *Handler) GetProject(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}
