package handlers

import (
	"fmt"
	"net/http"

	"lpmaarifnu-site-api/internal/models"
	"lpmaarifnu-site-api/internal/services"
	"lpmaarifnu-site-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type PengurusHandler struct {
	pengurusService services.PengurusService
}

func NewPengurusHandler(pengurusService services.PengurusService) *PengurusHandler {
	return &PengurusHandler{pengurusService: pengurusService}
}

// GetAllPengurus handles GET /api/v1/organization/pengurus
func (h *PengurusHandler) GetAllPengurus(c *gin.Context) {
	filters := make(map[string]interface{})

	if periode := c.Query("periode"); periode != "" {
		filters["periode"] = periode
	}

	if kategori := c.Query("kategori"); kategori != "" {
		filters["kategori"] = kategori
	}

	if active := c.Query("active"); active == "false" {
		filters["active"] = false
	} else {
		filters["active"] = true
	}

	pengurusList, err := h.pengurusService.GetAllPengurus(filters)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve pengurus", err.Error())
		return
	}

	// Format response
	var periode string
	if len(pengurusList) > 0 {
		periode = fmt.Sprintf("%d-%d", pengurusList[0].PeriodeMulai, pengurusList[0].PeriodeSelesai)
	} else {
		periode = c.Query("periode")
		if periode == "" {
			periode = "2024-2029"
		}
	}

	// Transform to response format
	var pengurusResponse []models.PengurusResponse
	for _, p := range pengurusList {
		resp := models.PengurusResponse{
			ID:       p.ID,
			Nama:     p.Nama,
			Jabatan:  p.Jabatan,
			Kategori: p.Kategori,
			Foto:     p.Foto,
			Bio:      p.Bio,
			Email:    p.Email,
			Phone:    p.Phone,
			Order:    p.OrderNumber,
			IsActive: p.IsActive,
		}
		resp.Periode.Mulai = p.PeriodeMulai
		resp.Periode.Selesai = p.PeriodeSelesai
		pengurusResponse = append(pengurusResponse, resp)
	}

	response := map[string]interface{}{
		"periode":  periode,
		"pengurus": pengurusResponse,
	}

	utils.SuccessResponse(c, http.StatusOK, "Pengurus retrieved successfully", response)
}
