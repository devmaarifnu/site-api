package repositories

import (
	"lpmaarifnu-site-api/internal/models"

	"gorm.io/gorm"
)

type NewsRepository interface {
	FindAll(offset, limit int, filters map[string]interface{}) ([]models.NewsArticle, int64, error)
	FindBySlug(slug string) (*models.NewsArticle, error)
	FindFeatured(limit int) ([]models.NewsArticle, error)
	IncrementViews(id uint) error
	GetRelatedArticles(categoryID uint, excludeID uint, limit int) ([]models.NewsArticle, error)
}

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &newsRepository{db: db}
}

func (r *newsRepository) FindAll(offset, limit int, filters map[string]interface{}) ([]models.NewsArticle, int64, error) {
	var articles []models.NewsArticle
	var total int64

	query := r.db.Model(&models.NewsArticle{}).
		Where("status = ?", "published").
		Where("deleted_at IS NULL")

	// Apply filters
	if categorySlug, ok := filters["category"].(string); ok && categorySlug != "" {
		query = query.Joins("JOIN categories ON categories.id = news_articles.category_id").
			Where("categories.slug = ?", categorySlug)
	}

	if search, ok := filters["search"].(string); ok && search != "" {
		query = query.Where("title LIKE ? OR excerpt LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if featured, ok := filters["featured"].(bool); ok {
		query = query.Where("is_featured = ?", featured)
	}

	// Count total
	query.Count(&total)

	// Get paginated results with preloading
	sort := "-published_at"
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
		Preload("Author").
		Preload("Tags").
		Offset(offset).
		Limit(limit).
		Find(&articles).Error

	return articles, total, err
}

func (r *newsRepository) FindBySlug(slug string) (*models.NewsArticle, error) {
	var article models.NewsArticle
	err := r.db.
		Preload("Category").
		Preload("Author").
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

func (r *newsRepository) FindFeatured(limit int) ([]models.NewsArticle, error) {
	var articles []models.NewsArticle
	err := r.db.
		Preload("Category").
		Preload("Author").
		Where("status = ?", "published").
		Where("is_featured = ?", true).
		Where("deleted_at IS NULL").
		Order("published_at DESC").
		Limit(limit).
		Find(&articles).Error

	return articles, err
}

func (r *newsRepository) IncrementViews(id uint) error {
	return r.db.Model(&models.NewsArticle{}).
		Where("id = ?", id).
		UpdateColumn("views", gorm.Expr("views + ?", 1)).Error
}

func (r *newsRepository) GetRelatedArticles(categoryID uint, excludeID uint, limit int) ([]models.NewsArticle, error) {
	var articles []models.NewsArticle
	err := r.db.
		Select("id, title, slug, image, published_at").
		Where("category_id = ?", categoryID).
		Where("id != ?", excludeID).
		Where("status = ?", "published").
		Where("deleted_at IS NULL").
		Order("published_at DESC").
		Limit(limit).
		Find(&articles).Error

	return articles, err
}
