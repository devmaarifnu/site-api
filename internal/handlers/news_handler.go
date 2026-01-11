package handlers

import (
	"fmt"
	"net/http"

	"lpmaarifnu-site-api/internal/services"
	"lpmaarifnu-site-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
	newsService services.NewsService
}

func NewNewsHandler(newsService services.NewsService) *NewsHandler {
	return &NewsHandler{newsService: newsService}
}

// GetAllNews handles GET /api/v1/news
func (h *NewsHandler) GetAllNews(c *gin.Context) {
	page, limit := utils.GetPaginationParams(c)
	offset := utils.CalculateOffset(page, limit)

	filters := make(map[string]interface{})
	if category := c.Query("category"); category != "" {
		filters["category"] = category
	}
	if search := c.Query("search"); search != "" {
		filters["search"] = search
	}
	if featured := c.Query("featured"); featured == "true" {
		filters["featured"] = true
	}
	if sort := c.Query("sort"); sort != "" {
		filters["sort"] = sort
	}

	articles, total, err := h.newsService.GetAllNews(offset, limit, filters)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve news articles", err.Error())
		return
	}

	response := utils.CreatePaginatedResponse(map[string]interface{}{
		"articles": articles,
	}, page, limit, total)

	utils.SuccessResponse(c, http.StatusOK, "News articles retrieved successfully", response)
}

// GetNewsBySlug handles GET /api/v1/news/:slug
func (h *NewsHandler) GetNewsBySlug(c *gin.Context) {
	slug := c.Param("slug")

	article, err := h.newsService.GetNewsBySlug(slug)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "News article not found", "Article with slug '"+slug+"' does not exist")
		return
	}

	// Increment views asynchronously
	go h.newsService.IncrementViews(article.ID)

	// Get related articles if category exists
	var relatedArticles interface{}
	if article.CategoryID != nil {
		related, _ := h.newsService.GetRelatedArticles(*article.CategoryID, article.ID, 3)
		relatedArticles = related
	}

	// Build response
	response := map[string]interface{}{
		"id":           article.ID,
		"title":        article.Title,
		"slug":         article.Slug,
		"excerpt":      article.Excerpt,
		"content":      article.Content,
		"image":        article.Image,
		"published_at": article.PublishedAt,
		"category":     article.Category,
		"author":       article.Author,
		"tags":         article.Tags,
		"views":        article.Views,
		"is_featured":  article.IsFeatured,
		"meta": map[string]string{
			"title":       article.MetaTitle,
			"description": article.MetaDescription,
			"keywords":    article.MetaKeywords,
		},
		"related_articles": relatedArticles,
		"created_at":       article.CreatedAt,
		"updated_at":       article.UpdatedAt,
	}

	utils.SuccessResponse(c, http.StatusOK, "News article retrieved successfully", response)
}

// GetFeaturedNews handles GET /api/v1/news/featured
func (h *NewsHandler) GetFeaturedNews(c *gin.Context) {
	limit := 5
	if l := c.Query("limit"); l != "" {
		if parsed, err := parseLimit(l); err == nil {
			limit = parsed
		}
	}

	articles, err := h.newsService.GetFeaturedNews(limit)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve featured news", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Featured news retrieved successfully", articles)
}

func parseLimit(s string) (int, error) {
	var limit int
	_, err := fmt.Sscanf(s, "%d", &limit)
	return limit, err
}
