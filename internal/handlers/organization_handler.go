package handlers

import (
	"net/http"

	"lpmaarifnu-site-api/internal/services"
	"lpmaarifnu-site-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	orgService services.OrganizationService
}

func NewOrganizationHandler(orgService services.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{orgService: orgService}
}

// GetOrganizationStructure handles GET /api/v1/organization/structure
func (h *OrganizationHandler) GetOrganizationStructure(c *gin.Context) {
	structure, err := h.orgService.GetOrganizationStructure()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve organization structure", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Organization structure retrieved successfully", structure)
}

// GetBoardMembers handles GET /api/v1/organization/board-members
func (h *OrganizationHandler) GetBoardMembers(c *gin.Context) {
	filters := make(map[string]interface{})
	if period := c.Query("period"); period != "" {
		filters["period"] = period
	}
	if active := c.Query("active"); active == "true" {
		filters["active"] = true
	} else if active == "false" {
		filters["active"] = false
	}

	members, err := h.orgService.GetBoardMembers(filters)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve board members", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Board members retrieved successfully", members)
}
