package repositories

import (
	"lpmaarifnu-site-api/internal/models"
	"time"

	"gorm.io/gorm"
)

type PengurusRepository interface {
	FindAll(filters map[string]interface{}) ([]models.Pengurus, error)
}

type pengurusRepository struct {
	db *gorm.DB
}

func NewPengurusRepository(db *gorm.DB) PengurusRepository {
	return &pengurusRepository{db: db}
}

func (r *pengurusRepository) FindAll(filters map[string]interface{}) ([]models.Pengurus, error) {
	var pengurusList []models.Pengurus
	query := r.db.Model(&models.Pengurus{})

	// Filter by period
	if periode, ok := filters["periode"].(string); ok && periode != "" {
		// Parse periode format: "2024-2029"
		// We'll check if current year is within the period
		query = query.Where("? BETWEEN periode_mulai AND periode_selesai", time.Now().Year())
	}

	// Filter by kategori
	if kategori, ok := filters["kategori"].(string); ok && kategori != "" {
		query = query.Where("kategori = ?", kategori)
	}

	// Filter by active status (default: true)
	if active, ok := filters["active"].(bool); ok {
		query = query.Where("is_active = ?", active)
	} else {
		query = query.Where("is_active = ?", true)
	}

	// Order by order_number
	query = query.Order("order_number ASC")

	err := query.Find(&pengurusList).Error
	return pengurusList, err
}
