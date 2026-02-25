package repositories

import (
	"lpmaarifnu-site-api/internal/models"

	"gorm.io/gorm"
)

type EditorialRepository interface {
	FindAllTeam() ([]models.EditorialTeam, error)
	FindAllCouncil() ([]models.EditorialCouncil, error)
}

type editorialRepository struct {
	db *gorm.DB
}

func NewEditorialRepository(db *gorm.DB) EditorialRepository {
	return &editorialRepository{db: db}
}

func (r *editorialRepository) FindAllTeam() ([]models.EditorialTeam, error) {
	var team []models.EditorialTeam
	err := r.db.
		Where("is_active = ?", true).
		Order("order_number ASC").
		Find(&team).Error
	return team, err
}

func (r *editorialRepository) FindAllCouncil() ([]models.EditorialCouncil, error) {
	var council []models.EditorialCouncil
	err := r.db.
		Where("is_active = ?", true).
		Order("order_number ASC").
		Find(&council).Error
	return council, err
}
