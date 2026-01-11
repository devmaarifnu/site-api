package models

import (
	"database/sql"
	"time"
)

type HeroSlide struct {
	ID                  uint         `gorm:"primaryKey" json:"id"`
	Title               string       `gorm:"size:500;not null" json:"title"`
	Description         string       `gorm:"type:text" json:"description,omitempty"`
	Image               string       `gorm:"size:500;not null" json:"image"`
	CTALabel            string       `gorm:"size:100" json:"cta_label,omitempty"`
	CTAHref             string       `gorm:"size:500" json:"cta_href,omitempty"`
	CTASecondaryLabel   string       `gorm:"size:100" json:"cta_secondary_label,omitempty"`
	CTASecondaryHref    string       `gorm:"size:500" json:"cta_secondary_href,omitempty"`
	OrderNumber         int          `gorm:"default:0" json:"order"`
	IsActive            bool         `gorm:"default:true" json:"is_active"`
	StartDate           sql.NullTime `json:"start_date,omitempty"`
	EndDate             sql.NullTime `json:"end_date,omitempty"`
	CreatedAt           time.Time    `json:"created_at"`
	UpdatedAt           time.Time    `json:"updated_at"`
}

func (HeroSlide) TableName() string {
	return "hero_slides"
}
