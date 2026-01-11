package repositories

import (
	"lpmaarifnu-site-api/internal/models"

	"gorm.io/gorm"
)

type HeroRepository interface {
	FindActiveSlides() ([]models.HeroSlide, error)
}

type heroRepository struct {
	db *gorm.DB
}

func NewHeroRepository(db *gorm.DB) HeroRepository {
	return &heroRepository{db: db}
}

func (r *heroRepository) FindActiveSlides() ([]models.HeroSlide, error) {
	var slides []models.HeroSlide
	err := r.db.
		Where("is_active = ?", true).
		Order("order_number ASC").
		Find(&slides).Error

	return slides, err
}
