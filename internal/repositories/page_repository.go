package repositories

import (
	"lpmaarifnu-site-api/internal/models"

	"gorm.io/gorm"
)

type PageRepository interface {
	FindBySlug(slug string) (*models.Page, error)
}

type pageRepository struct {
	db *gorm.DB
}

func NewPageRepository(db *gorm.DB) PageRepository {
	return &pageRepository{db: db}
}

func (r *pageRepository) FindBySlug(slug string) (*models.Page, error) {
	var page models.Page
	err := r.db.
		Where("slug = ?", slug).
		Where("is_active = ?", true).
		First(&page).Error

	if err != nil {
		return nil, err
	}

	return &page, nil
}
