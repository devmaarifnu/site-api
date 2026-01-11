package repositories

import (
	"lpmaarifnu-site-api/internal/models"

	"gorm.io/gorm"
)

type OpinionRepository interface {
	FindAll(offset, limit int, filters map[string]interface{}) ([]models.OpinionArticle, int64, error)
	FindBySlug(slug string) (*models.OpinionArticle, error)
	IncrementViews(id uint) error
}

type opinionRepository struct {
	db *gorm.DB
}

func NewOpinionRepository(db *gorm.DB) OpinionRepository {
	return &opinionRepository{db: db}
}

func (r *opinionRepository) FindAll(offset, limit int, filters map[string]interface{}) ([]models.OpinionArticle, int64, error) {
	var articles []models.OpinionArticle
	var total int64

	query := r.db.Model(&models.OpinionArticle{}).
		Where("status = ?", "published").
		Where("deleted_at IS NULL")

	// Apply filters
	if search, ok := filters["search"].(string); ok && search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	// Count total
	query.Count(&total)

	// Get paginated results
	err := query.
		Preload("Tags").
		Order("published_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&articles).Error

	return articles, total, err
}

func (r *opinionRepository) FindBySlug(slug string) (*models.OpinionArticle, error) {
	var article models.OpinionArticle
	err := r.db.
		Preload("Tags").
		Where("slug = ?", slug).
		Where("status = ?", "published").
		Where("deleted_at IS NULL").
		First(&article).Error

	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (r *opinionRepository) IncrementViews(id uint) error {
	return r.db.Model(&models.OpinionArticle{}).
		Where("id = ?", id).
		UpdateColumn("views", gorm.Expr("views + ?", 1)).Error
}
