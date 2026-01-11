package models

import "time"

type Category struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Slug        string    `gorm:"size:100;unique;not null" json:"slug"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	Type        string    `gorm:"type:enum('news','opinion','document');not null" json:"type"`
	Icon        string    `gorm:"size:100" json:"icon,omitempty"`
	Color       string    `gorm:"size:7" json:"color,omitempty"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	OrderNumber int       `gorm:"default:0" json:"order_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Category) TableName() string {
	return "categories"
}
