package repositories

import (
	"lpmaarifnu-site-api/internal/models"

	"gorm.io/gorm"
)

type TagRepository interface {
	FindAll(limit int) ([]models.Tag, error)
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) FindAll(limit int) ([]models.Tag, error) {
	var tags []models.Tag

	query := r.db.Model(&models.Tag{})

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Order("name ASC").Find(&tags).Error
	return tags, err
}
