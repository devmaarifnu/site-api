package repositories

import (
	"lpmaarifnu-site-api/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll(filters map[string]interface{}) ([]models.Category, error)
	FindBySlug(slug string) (*models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) FindAll(filters map[string]interface{}) ([]models.Category, error) {
	var categories []models.Category

	query := r.db.Model(&models.Category{}).Where("is_active = ?", true)

	if categoryType, ok := filters["type"].(string); ok && categoryType != "" {
		query = query.Where("type = ?", categoryType)
	}

	err := query.Order("order_number ASC").Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) FindBySlug(slug string) (*models.Category, error) {
	var category models.Category
	err := r.db.
		Where("slug = ?", slug).
		Where("is_active = ?", true).
		First(&category).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}
