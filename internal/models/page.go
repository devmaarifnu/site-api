package models

import "time"

type Page struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Slug            string    `gorm:"size:255;unique;not null" json:"slug"`
	Title           string    `gorm:"size:500;not null" json:"title"`
	Content         string    `gorm:"type:longtext" json:"content,omitempty"`
	Metadata        string    `gorm:"type:json" json:"metadata,omitempty"`
	Template        string    `gorm:"size:100;default:'default'" json:"template"`
	IsActive        bool      `gorm:"default:true" json:"is_active"`
	MetaTitle       string    `gorm:"size:255" json:"meta_title,omitempty"`
	MetaDescription string    `gorm:"type:text" json:"meta_description,omitempty"`
	MetaKeywords    string    `gorm:"size:500" json:"meta_keywords,omitempty"`
	LastUpdatedBy   *uint     `json:"last_updated_by,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (Page) TableName() string {
	return "pages"
}
