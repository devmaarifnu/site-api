package models

import (
	"database/sql"
	"time"
)

type Document struct {
	ID            uint         `gorm:"primaryKey" json:"id"`
	Title         string       `gorm:"size:500;not null" json:"title"`
	Description   string       `gorm:"type:text" json:"description,omitempty"`
	CategoryID    *uint        `json:"category_id,omitempty"`
	Category      *Category    `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	FileName      string       `gorm:"size:255;not null" json:"file_name"`
	FilePath      string       `gorm:"size:500;not null" json:"file_path"`
	FileType      string       `gorm:"size:50;not null" json:"file_type"`
	FileSize      uint64       `json:"file_size"`
	MimeType      string       `gorm:"size:100" json:"mime_type,omitempty"`
	DownloadCount uint         `gorm:"default:0" json:"download_count"`
	IsPublic      bool         `gorm:"default:true" json:"is_public"`
	UploadedBy    *uint        `json:"uploaded_by,omitempty"`
	Status        string       `gorm:"type:enum('active','archived');default:'active'" json:"status"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	DeletedAt     sql.NullTime `gorm:"index" json:"deleted_at,omitempty"`
}

func (Document) TableName() string {
	return "documents"
}
