package handlers

import (
	"net/http"
	"strconv"

	"lpmaarifnu-site-api/internal/models"
	"lpmaarifnu-site-api/internal/services"
	"lpmaarifnu-site-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	eventService services.EventService
}

func NewEventHandler(eventService services.EventService) *EventHandler {
	return &EventHandler{eventService: eventService}
}

// GetAllEvents handles GET /api/v1/events/flayers
func (h *EventHandler) GetAllEvents(c *gin.Context) {
	filters := make(map[string]interface{})

	// Parse limit
	limit := 10
	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 100 {
			limit = parsed
		}
	}
	filters["limit"] = limit

	// Parse active filter
	if active := c.Query("active"); active == "false" {
		filters["active"] = false
	} else {
		filters["active"] = true
	}

	// Parse upcoming filter
	if upcoming := c.Query("upcoming"); upcoming == "true" {
		filters["upcoming"] = true
	}

	// Parse check_display_period filter (optional, default: false)
	if checkDisplayPeriod := c.Query("check_display_period"); checkDisplayPeriod == "true" {
		filters["check_display_period"] = true
	}

	events, err := h.eventService.GetAllEvents(filters)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve event flayers", err.Error())
		return
	}

	// Transform to response format
	var eventsResponse []models.EventFlayerResponse
	for _, e := range events {
		resp := models.EventFlayerResponse{
			ID:          e.ID,
			Title:       e.Title,
			Image:       e.Image,
			Order:       e.OrderNumber,
			IsActive:    e.IsActive,
		}

		if e.Description != nil {
			resp.Description = *e.Description
		}

		if e.EventDate != nil && e.EventDate.Valid {
			resp.EventDate = e.EventDate.Time.Format("2006-01-02")
		}

		if e.EventLocation != nil {
			resp.EventLocation = *e.EventLocation
		}

		if e.RegistrationURL != nil {
			resp.RegistrationURL = *e.RegistrationURL
		}

		// Build contact info
		if e.ContactPerson != nil || e.ContactPhone != nil || e.ContactEmail != nil {
			resp.Contact = &struct {
				Person string `json:"person,omitempty"`
				Phone  string `json:"phone,omitempty"`
				Email  string `json:"email,omitempty"`
			}{}

			if e.ContactPerson != nil {
				resp.Contact.Person = *e.ContactPerson
			}
			if e.ContactPhone != nil {
				resp.Contact.Phone = *e.ContactPhone
			}
			if e.ContactEmail != nil {
				resp.Contact.Email = *e.ContactEmail
			}
		}

		// Build display period
		if e.StartDisplayDate != nil || e.EndDisplayDate != nil {
			resp.DisplayPeriod = &struct {
				Start string `json:"start,omitempty"`
				End   string `json:"end,omitempty"`
			}{}

			if e.StartDisplayDate != nil {
				resp.DisplayPeriod.Start = e.StartDisplayDate.Format("2006-01-02T15:04:05Z07:00")
			}
			if e.EndDisplayDate != nil {
				resp.DisplayPeriod.End = e.EndDisplayDate.Format("2006-01-02T15:04:05Z07:00")
			}
		}

		eventsResponse = append(eventsResponse, resp)
	}

	utils.SuccessResponse(c, http.StatusOK, "Event flayers retrieved successfully", eventsResponse)
}
