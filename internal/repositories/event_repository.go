package repositories

import (
	"lpmaarifnu-site-api/internal/models"
	"time"

	"gorm.io/gorm"
)

type EventRepository interface {
	FindAll(filters map[string]interface{}) ([]models.EventFlayer, error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) FindAll(filters map[string]interface{}) ([]models.EventFlayer, error) {
	var events []models.EventFlayer
	query := r.db.Model(&models.EventFlayer{})

	// Filter by active status (default: true)
	if active, ok := filters["active"].(bool); ok {
		query = query.Where("is_active = ?", active)
	} else {
		query = query.Where("is_active = ?", true)
	}

	// Filter by display date range (only if check_display_period is true)
	if checkDisplayPeriod, ok := filters["check_display_period"].(bool); ok && checkDisplayPeriod {
		now := time.Now()
		query = query.Where("(start_display_date IS NULL OR start_display_date <= ?)", now)
		query = query.Where("(end_display_date IS NULL OR end_display_date >= ?)", now)
	}

	// Filter upcoming events only
	if upcoming, ok := filters["upcoming"].(bool); ok && upcoming {
		now := time.Now()
		query = query.Where("event_date >= ?", now)
	}

	// Apply limit
	limit := 10
	if l, ok := filters["limit"].(int); ok && l > 0 && l <= 100 {
		limit = l
	}

	// Order by order_number ASC, then event_date DESC
	err := query.
		Order("order_number ASC, event_date DESC").
		Limit(limit).
		Find(&events).Error

	return events, err
}
