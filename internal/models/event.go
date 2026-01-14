package models

import (
	"database/sql"
	"time"
)

// EventFlayer represents an event/activity flyer
type EventFlayer struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	Title            string         `json:"title" gorm:"size:500;not null"`
	Description      *string        `json:"description" gorm:"type:text"`
	Image            string         `json:"image" gorm:"size:500;not null"`
	EventDate        *sql.NullTime  `json:"event_date" gorm:"type:date"`
	EventLocation    *string        `json:"event_location" gorm:"size:500"`
	RegistrationURL  *string        `json:"registration_url" gorm:"size:500"`
	ContactPerson    *string        `json:"contact_person" gorm:"size:255"`
	ContactPhone     *string        `json:"contact_phone" gorm:"size:20"`
	ContactEmail     *string        `json:"contact_email" gorm:"size:255"`
	OrderNumber      int            `json:"order_number" gorm:"default:0"`
	IsActive         bool           `json:"is_active" gorm:"default:true"`
	StartDisplayDate *time.Time     `json:"start_display_date"`
	EndDisplayDate   *time.Time     `json:"end_display_date"`
	CreatedBy        *uint          `json:"created_by"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}

// TableName overrides the table name
func (EventFlayer) TableName() string {
	return "event_flayers"
}

// EventFlayerResponse represents the formatted response
type EventFlayerResponse struct {
	ID              uint   `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description,omitempty"`
	Image           string `json:"image"`
	EventDate       string `json:"event_date,omitempty"`
	EventLocation   string `json:"event_location,omitempty"`
	RegistrationURL string `json:"registration_url,omitempty"`
	Contact         *struct {
		Person string `json:"person,omitempty"`
		Phone  string `json:"phone,omitempty"`
		Email  string `json:"email,omitempty"`
	} `json:"contact,omitempty"`
	DisplayPeriod *struct {
		Start string `json:"start,omitempty"`
		End   string `json:"end,omitempty"`
	} `json:"display_period,omitempty"`
	Order    int  `json:"order"`
	IsActive bool `json:"is_active"`
}
