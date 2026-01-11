package services

import (
	"lpmaarifnu-site-api/internal/models"
	"lpmaarifnu-site-api/internal/repositories"
)

type NewsService interface {
	GetAllNews(offset, limit int, filters map[string]interface{}) ([]models.NewsArticle, int64, error)
	GetNewsBySlug(slug string) (*models.NewsArticle, error)
	GetFeaturedNews(limit int) ([]models.NewsArticle, error)
	IncrementViews(id uint) error
	GetRelatedArticles(categoryID uint, excludeID uint, limit int) ([]models.NewsArticle, error)
}

type newsService struct {
	newsRepo repositories.NewsRepository
}

func NewNewsService(newsRepo repositories.NewsRepository) NewsService {
	return &newsService{newsRepo: newsRepo}
}

func (s *newsService) GetAllNews(offset, limit int, filters map[string]interface{}) ([]models.NewsArticle, int64, error) {
	return s.newsRepo.FindAll(offset, limit, filters)
}

func (s *newsService) GetNewsBySlug(slug string) (*models.NewsArticle, error) {
	return s.newsRepo.FindBySlug(slug)
}

func (s *newsService) GetFeaturedNews(limit int) ([]models.NewsArticle, error) {
	return s.newsRepo.FindFeatured(limit)
}

func (s *newsService) IncrementViews(id uint) error {
	return s.newsRepo.IncrementViews(id)
}

func (s *newsService) GetRelatedArticles(categoryID uint, excludeID uint, limit int) ([]models.NewsArticle, error) {
	return s.newsRepo.GetRelatedArticles(categoryID, excludeID, limit)
}
