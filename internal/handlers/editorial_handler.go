package handlers

import (
	"net/http"

	"lpmaarifnu-site-api/internal/services"
	"lpmaarifnu-site-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type EditorialHandler struct {
	editorialService services.EditorialService
}

func NewEditorialHandler(editorialService services.EditorialService) *EditorialHandler {
	return &EditorialHandler{editorialService: editorialService}
}

// GetEditorialTeam handles GET /api/v1/editorial/team
func (h *EditorialHandler) GetEditorialTeam(c *gin.Context) {
	team, err := h.editorialService.GetEditorialTeam()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve editorial team", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Editorial team retrieved successfully", team)
}
