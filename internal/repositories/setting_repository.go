package repositories

import (
	"lpmaarifnu-site-api/internal/models"

	"gorm.io/gorm"
)

type SettingRepository interface {
	FindPublicSettings() ([]models.Setting, error)
}

type settingRepository struct {
	db *gorm.DB
}

func NewSettingRepository(db *gorm.DB) SettingRepository {
	return &settingRepository{db: db}
}

func (r *settingRepository) FindPublicSettings() ([]models.Setting, error) {
	var settings []models.Setting
	err := r.db.
		Where("is_public = ?", true).
		Find(&settings).Error

	return settings, err
}
