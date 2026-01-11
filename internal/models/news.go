package models

import (
	"database/sql"
	"time"
)

type NewsArticle struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Title           string         `gorm:"size:500;not null" json:"title"`
	Slug            string         `gorm:"size:500;unique;not null" json:"slug"`
	Excerpt         string         `gorm:"type:text;not null" json:"excerpt"`
	Content         string         `gorm:"type:longtext;not null" json:"content"`
	Image           string         `gorm:"size:500" json:"image,omitempty"`
	CategoryID      *uint          `json:"category_id,omitempty"`
	Category        *Category      `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	AuthorID        *uint          `json:"author_id,omitempty"`
	Author          *User          `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Status          string         `gorm:"type:enum('draft','published','archived');default:'draft'" json:"status"`
	PublishedAt     sql.NullTime   `json:"published_at,omitempty"`
	Views           uint           `gorm:"default:0" json:"views"`
	IsFeatured      bool           `gorm:"default:false" json:"is_featured"`
	MetaTitle       string         `gorm:"size:255" json:"meta_title,omitempty"`
	MetaDescription string         `gorm:"type:text" json:"meta_description,omitempty"`
	MetaKeywords    string         `gorm:"size:500" json:"meta_keywords,omitempty"`
	Tags            []Tag          `gorm:"many2many:news_tags;" json:"tags,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       sql.NullTime   `gorm:"index" json:"deleted_at,omitempty"`
}

func (NewsArticle) TableName() string {
	return "news_articles"
}

type User struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"size:255;not null" json:"name"`
	Email  string `gorm:"size:255;unique;not null" json:"email,omitempty"`
	Avatar string `gorm:"size:500" json:"avatar,omitempty"`
}

func (User) TableName() string {
	return "users"
}
