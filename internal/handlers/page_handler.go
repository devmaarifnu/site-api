package handlers

import (
	"net/http"

	"lpmaarifnu-site-api/internal/services"
	"lpmaarifnu-site-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type PageHandler struct {
	pageService services.PageService
}

func NewPageHandler(pageService services.PageService) *PageHandler {
	return &PageHandler{pageService: pageService}
}

// GetPageBySlug handles GET /api/v1/pages/:slug
func (h *PageHandler) GetPageBySlug(c *gin.Context) {
	slug := c.Param("slug")

	page, err := h.pageService.GetPageBySlug(slug)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Page not found", "Page with slug '"+slug+"' does not exist")
		return
	}

	response := map[string]interface{}{
		"slug":     page.Slug,
		"title":    page.Title,
		"template": page.Template,
		"content":  page.Metadata, // JSON metadata contains the actual content structure
		"meta": map[string]string{
			"title":       page.MetaTitle,
			"description": page.MetaDescription,
			"keywords":    page.MetaKeywords,
		},
		"is_active":  page.IsActive,
		"updated_at": page.UpdatedAt,
	}

	utils.SuccessResponse(c, http.StatusOK, "Page content retrieved successfully", response)
}
