package models

import (
	"database/sql"
	"time"
)

type OpinionArticle struct {
	ID              uint         `gorm:"primaryKey" json:"id"`
	Title           string       `gorm:"size:500;not null" json:"title"`
	Slug            string       `gorm:"size:500;unique;not null" json:"slug"`
	Excerpt         string       `gorm:"type:text;not null" json:"excerpt"`
	Content         string       `gorm:"type:longtext;not null" json:"content"`
	Image           string       `gorm:"size:500" json:"image,omitempty"`
	AuthorName      string       `gorm:"size:255;not null" json:"author_name"`
	AuthorTitle     string       `gorm:"size:255" json:"author_title,omitempty"`
	AuthorImage     string       `gorm:"size:500" json:"author_image,omitempty"`
	AuthorBio       string       `gorm:"type:text" json:"author_bio,omitempty"`
	Status          string       `gorm:"type:enum('draft','published','archived');default:'draft'" json:"status"`
	PublishedAt     sql.NullTime `json:"published_at,omitempty"`
	Views           uint         `gorm:"default:0" json:"views"`
	IsFeatured      bool         `gorm:"default:false" json:"is_featured"`
	MetaTitle       string       `gorm:"size:255" json:"meta_title,omitempty"`
	MetaDescription string       `gorm:"type:text" json:"meta_description,omitempty"`
	MetaKeywords    string       `gorm:"size:500" json:"meta_keywords,omitempty"`
	CreatedBy       *uint        `json:"created_by,omitempty"`
	Tags            []Tag        `gorm:"many2many:opinion_tags;" json:"tags,omitempty"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	DeletedAt       sql.NullTime `gorm:"index" json:"deleted_at,omitempty"`
}

func (OpinionArticle) TableName() string {
	return "opinion_articles"
}
