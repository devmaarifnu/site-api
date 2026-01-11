package repositories

import (
	"lpmaarifnu-site-api/internal/models"

	"gorm.io/gorm"
)

type DocumentRepository interface {
	FindAll(offset, limit int, filters map[string]interface{}) ([]models.Document, int64, error)
	FindByID(id uint) (*models.Document, error)
	IncrementDownloads(id uint) error
}

type documentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) DocumentRepository {
	return &documentRepository{db: db}
}

func (r *documentRepository) FindAll(offset, limit int, filters map[string]interface{}) ([]models.Document, int64, error) {
	var documents []models.Document
	var total int64

	query := r.db.Model(&models.Document{}).
		Where("status = ?", "active").
		Where("is_public = ?", true).
		Where("deleted_at IS NULL")

	// Apply filters
	if categorySlug, ok := filters["category"].(string); ok && categorySlug != "" {
		query = query.Joins("JOIN categories ON categories.id = documents.category_id").
			Where("categories.slug = ?", categorySlug)
	}

	if search, ok := filters["search"].(string); ok && search != "" {
		query = query.Where("title LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Count total
	query.Count(&total)

	// Get paginated results
	sort := "-created_at"
	if s, ok := filters["sort"].(string); ok && s != "" {
		sort = s
	}

	if sort[0] == '-' {
		query = query.Order(sort[1:] + " DESC")
	} else {
		query = query.Order(sort + " ASC")
	}

	err := query.
		Preload("Category").
		Offset(offset).
		Limit(limit).
		Find(&documents).Error

	return documents, total, err
}

func (r *documentRepository) FindByID(id uint) (*models.Document, error) {
	var document models.Document
	err := r.db.
		Preload("Category").
		Where("id = ?", id).
		Where("status = ?", "active").
		Where("is_public = ?", true).
		Where("deleted_at IS NULL").
		First(&document).Error

	if err != nil {
		return nil, err
	}

	return &document, nil
}

func (r *documentRepository) IncrementDownloads(id uint) error {
	return r.db.Model(&models.Document{}).
		Where("id = ?", id).
		UpdateColumn("download_count", gorm.Expr("download_count + ?", 1)).Error
}
