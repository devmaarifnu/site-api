package handlers

import (
	"net/http"

	"lpmaarifnu-site-api/internal/services"
	"lpmaarifnu-site-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type SettingHandler struct {
	settingService services.SettingService
}

func NewSettingHandler(settingService services.SettingService) *SettingHandler {
	return &SettingHandler{settingService: settingService}
}

// GetPublicSettings handles GET /api/v1/settings
func (h *SettingHandler) GetPublicSettings(c *gin.Context) {
	settings, err := h.settingService.GetPublicSettings()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve settings", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Settings retrieved successfully", settings)
}
