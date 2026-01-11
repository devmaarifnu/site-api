package utils

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	CurrentPage   int  `json:"current_page"`
	TotalPages    int  `json:"total_pages"`
	TotalItems    int64 `json:"total_items"`
	ItemsPerPage  int  `json:"items_per_page"`
	HasNext       bool `json:"has_next"`
	HasPrev       bool `json:"has_prev"`
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

// GetPaginationParams extracts pagination parameters from request
func GetPaginationParams(c *gin.Context) (page int, limit int) {
	page = 1
	limit = 10

	if p := c.Query("page"); p != "" {
		if pageNum, err := strconv.Atoi(p); err == nil && pageNum > 0 {
			page = pageNum
		}
	}

	if l := c.Query("limit"); l != "" {
		if limitNum, err := strconv.Atoi(l); err == nil && limitNum > 0 && limitNum <= 100 {
			limit = limitNum
		}
	}

	return page, limit
}

// CalculateOffset calculates database offset from page and limit
func CalculateOffset(page, limit int) int {
	return (page - 1) * limit
}

// CreatePagination creates pagination metadata
func CreatePagination(page, limit int, total int64) Pagination {
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return Pagination{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalItems:   total,
		ItemsPerPage: limit,
		HasNext:      page < totalPages,
		HasPrev:      page > 1,
	}
}

// CreatePaginatedResponse creates a paginated response
func CreatePaginatedResponse(data interface{}, page, limit int, total int64) PaginatedResponse {
	return PaginatedResponse{
		Data:       data,
		Pagination: CreatePagination(page, limit, total),
	}
}
