package handlers

import (
	"net/http"

	"lpmaarifnu-site-api/internal/services"
	"lpmaarifnu-site-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService services.CategoryService
	newsService     services.NewsService
}

func NewCategoryHandler(categoryService services.CategoryService, newsService services.NewsService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
		newsService:     newsService,
	}
}

// GetAllCategories handles GET /api/v1/categories
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	filters := make(map[string]interface{})
	if categoryType := c.Query("type"); categoryType != "" {
		filters["type"] = categoryType
	}

	categories, err := h.categoryService.GetAllCategories(filters)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve categories", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Categories retrieved successfully", categories)
}

// GetCategoryBySlug handles GET /api/v1/categories/:slug
func (h *CategoryHandler) GetCategoryBySlug(c *gin.Context) {
	slug := c.Param("slug")

	category, err := h.categoryService.GetCategoryBySlug(slug)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Category not found", "Category with slug '"+slug+"' does not exist")
		return
	}

	// Get latest articles in this category
	filters := map[string]interface{}{
		"category": slug,
	}
	latestArticles, _, _ := h.newsService.GetAllNews(0, 5, filters)

	response := map[string]interface{}{
		"id":              category.ID,
		"name":            category.Name,
		"slug":            category.Slug,
		"description":     category.Description,
		"type":            category.Type,
		"icon":            category.Icon,
		"color":           category.Color,
		"latest_articles": latestArticles,
	}

	utils.SuccessResponse(c, http.StatusOK, "Category retrieved successfully", response)
}
