package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"lpmaarifnu-site-api/internal/services"
	"lpmaarifnu-site-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type DocumentHandler struct {
	documentService services.DocumentService
}

func NewDocumentHandler(documentService services.DocumentService) *DocumentHandler {
	return &DocumentHandler{documentService: documentService}
}

// GetAllDocuments handles GET /api/v1/documents
func (h *DocumentHandler) GetAllDocuments(c *gin.Context) {
	page, limit := utils.GetPaginationParams(c)
	offset := utils.CalculateOffset(page, limit)

	filters := make(map[string]interface{})
	if category := c.Query("category"); category != "" {
		filters["category"] = category
	}
	if search := c.Query("search"); search != "" {
		filters["search"] = search
	}
	if sort := c.Query("sort"); sort != "" {
		filters["sort"] = sort
	}

	documents, total, err := h.documentService.GetAllDocuments(offset, limit, filters)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve documents", err.Error())
		return
	}

	// Format file size
	formattedDocs := make([]map[string]interface{}, len(documents))
	for i, doc := range documents {
		formattedDocs[i] = map[string]interface{}{
			"id":                  doc.ID,
			"title":               doc.Title,
			"description":         doc.Description,
			"category":            doc.Category,
			"file_name":           doc.FileName,
			"file_type":           doc.FileType,
			"file_size":           doc.FileSize,
			"file_size_formatted": formatFileSize(doc.FileSize),
			"download_url":        doc.FilePath,
			"download_count":      doc.DownloadCount,
			"uploaded_at":         doc.CreatedAt,
			"created_at":          doc.CreatedAt,
		}
	}

	response := utils.CreatePaginatedResponse(map[string]interface{}{
		"documents": formattedDocs,
	}, page, limit, total)

	utils.SuccessResponse(c, http.StatusOK, "Documents retrieved successfully", response)
}

// GetDocumentByID handles GET /api/v1/documents/:id
func (h *DocumentHandler) GetDocumentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid document ID", err.Error())
		return
	}

	document, err := h.documentService.GetDocumentByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Document not found", "Document with ID "+idStr+" does not exist")
		return
	}

	// Increment downloads asynchronously
	go h.documentService.IncrementDownloads(document.ID)

	response := map[string]interface{}{
		"id":                  document.ID,
		"title":               document.Title,
		"description":         document.Description,
		"category":            document.Category,
		"file_name":           document.FileName,
		"file_type":           document.FileType,
		"file_size":           document.FileSize,
		"file_size_formatted": formatFileSize(document.FileSize),
		"mime_type":           document.MimeType,
		"download_url":        document.FilePath,
		"download_count":      document.DownloadCount,
		"uploaded_at":         document.CreatedAt,
		"created_at":          document.CreatedAt,
		"updated_at":          document.UpdatedAt,
	}

	utils.SuccessResponse(c, http.StatusOK, "Document retrieved successfully", response)
}

func formatFileSize(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
