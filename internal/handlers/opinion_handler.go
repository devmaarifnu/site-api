package handlers

import (
	"net/http"

	"lpmaarifnu-site-api/internal/services"
	"lpmaarifnu-site-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type OpinionHandler struct {
	opinionService services.OpinionService
}

func NewOpinionHandler(opinionService services.OpinionService) *OpinionHandler {
	return &OpinionHandler{opinionService: opinionService}
}

// GetAllOpinions handles GET /api/v1/opinions
func (h *OpinionHandler) GetAllOpinions(c *gin.Context) {
	page, limit := utils.GetPaginationParams(c)
	offset := utils.CalculateOffset(page, limit)

	filters := make(map[string]interface{})
	if search := c.Query("search"); search != "" {
		filters["search"] = search
	}

	articles, total, err := h.opinionService.GetAllOpinions(offset, limit, filters)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve opinion articles", err.Error())
		return
	}

	response := utils.CreatePaginatedResponse(map[string]interface{}{
		"articles": articles,
	}, page, limit, total)

	utils.SuccessResponse(c, http.StatusOK, "Opinion articles retrieved successfully", response)
}

// GetOpinionBySlug handles GET /api/v1/opinions/:slug
func (h *OpinionHandler) GetOpinionBySlug(c *gin.Context) {
	slug := c.Param("slug")

	article, err := h.opinionService.GetOpinionBySlug(slug)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Opinion article not found", "Article with slug '"+slug+"' does not exist")
		return
	}

	// Increment views asynchronously
	go h.opinionService.IncrementViews(article.ID)

	// Build response with author details
	response := map[string]interface{}{
		"id":           article.ID,
		"title":        article.Title,
		"slug":         article.Slug,
		"excerpt":      article.Excerpt,
		"content":      article.Content,
		"image":        article.Image,
		"published_at": article.PublishedAt,
		"author": map[string]string{
			"name":  article.AuthorName,
			"title": article.AuthorTitle,
			"image": article.AuthorImage,
			"bio":   article.AuthorBio,
		},
		"tags":        article.Tags,
		"views":       article.Views,
		"is_featured": article.IsFeatured,
		"meta": map[string]string{
			"title":       article.MetaTitle,
			"description": article.MetaDescription,
			"keywords":    article.MetaKeywords,
		},
		"created_at": article.CreatedAt,
		"updated_at": article.UpdatedAt,
	}

	utils.SuccessResponse(c, http.StatusOK, "Opinion article retrieved successfully", response)
}
