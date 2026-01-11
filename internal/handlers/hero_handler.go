package handlers

import (
	"net/http"

	"lpmaarifnu-site-api/internal/services"
	"lpmaarifnu-site-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type HeroHandler struct {
	heroService services.HeroService
}

func NewHeroHandler(heroService services.HeroService) *HeroHandler {
	return &HeroHandler{heroService: heroService}
}

// GetActiveSlides handles GET /api/v1/hero-slides
func (h *HeroHandler) GetActiveSlides(c *gin.Context) {
	slides, err := h.heroService.GetActiveSlides()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve hero slides", err.Error())
		return
	}

	// Format response with CTA structure
	formattedSlides := make([]map[string]interface{}, len(slides))
	for i, slide := range slides {
		cta := map[string]interface{}{}
		if slide.CTALabel != "" {
			cta["label"] = slide.CTALabel
			cta["href"] = slide.CTAHref
		}
		if slide.CTASecondaryLabel != "" {
			cta["secondary"] = map[string]string{
				"label": slide.CTASecondaryLabel,
				"href":  slide.CTASecondaryHref,
			}
		}

		formattedSlides[i] = map[string]interface{}{
			"id":          slide.ID,
			"title":       slide.Title,
			"description": slide.Description,
			"image":       slide.Image,
			"cta":         cta,
			"order":       slide.OrderNumber,
			"is_active":   slide.IsActive,
			"created_at":  slide.CreatedAt,
		}
	}

	utils.SuccessResponse(c, http.StatusOK, "Hero slides retrieved successfully", formattedSlides)
}
