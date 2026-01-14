package handlers

import (
	"net/http"

	"lpmaarifnu-site-api/internal/models"
	"lpmaarifnu-site-api/internal/services"
	"lpmaarifnu-site-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	contactService services.ContactService
}

func NewContactHandler(contactService services.ContactService) *ContactHandler {
	return &ContactHandler{contactService: contactService}
}

// SubmitContact handles POST /api/v1/contact/submit
func (h *ContactHandler) SubmitContact(c *gin.Context) {
	var req models.ContactRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", err.Error())
		return
	}

	// Get IP address and user agent
	ipAddress := c.ClientIP()
	userAgent := c.Request.UserAgent()

	contact, err := h.contactService.SubmitContactMessage(&req, ipAddress, userAgent)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to submit message", err.Error())
		return
	}

	response := map[string]interface{}{
		"ticket_id":  contact.TicketID,
		"name":       contact.Name,
		"email":      contact.Email,
		"subject":    contact.Subject,
		"status":     contact.Status,
		"created_at": contact.CreatedAt,
	}

	utils.SuccessResponse(c, http.StatusCreated, "Pesan Anda telah terkirim. Kami akan segera menghubungi Anda.", response)
}
