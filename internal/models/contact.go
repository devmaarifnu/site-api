package models

import (
	"time"
)

// ContactMessage represents a contact form submission
type ContactMessage struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	TicketID   string     `json:"ticket_id" gorm:"size:50;unique;not null"`
	Name       string     `json:"name" gorm:"size:255;not null"`
	Email      string     `json:"email" gorm:"size:255;not null"`
	Phone      *string    `json:"phone" gorm:"size:20"`
	Subject    string     `json:"subject" gorm:"size:500;not null"`
	Message    string     `json:"message" gorm:"type:text;not null"`
	Status     string     `json:"status" gorm:"type:enum('new','read','in_progress','resolved','closed');default:'new'"`
	Priority   string     `json:"priority" gorm:"type:enum('low','medium','high','urgent');default:'medium'"`
	IPAddress  *string    `json:"ip_address,omitempty" gorm:"size:45"`
	UserAgent  *string    `json:"user_agent,omitempty" gorm:"type:text"`
	AssignedTo *uint      `json:"assigned_to,omitempty"`
	RepliedAt  *time.Time `json:"replied_at,omitempty"`
	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
	Notes      *string    `json:"notes,omitempty" gorm:"type:text"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// TableName overrides the table name
func (ContactMessage) TableName() string {
	return "contact_messages"
}

// ContactRequest represents the request body for submitting a contact message
type ContactRequest struct {
	Name    string  `json:"name" binding:"required,min=3,max=255"`
	Email   string  `json:"email" binding:"required,email"`
	Phone   *string `json:"phone" binding:"omitempty,max=20"`
	Subject string  `json:"subject" binding:"required,min=5,max=500"`
	Message string  `json:"message" binding:"required,min=10"`
}
