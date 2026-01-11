package services

import (
	"lpmaarifnu-site-api/internal/models"
	"lpmaarifnu-site-api/internal/repositories"
)

type CategoryService interface {
	GetAllCategories(filters map[string]interface{}) ([]models.Category, error)
	GetCategoryBySlug(slug string) (*models.Category, error)
}

type categoryService struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryService(categoryRepo repositories.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (s *categoryService) GetAllCategories(filters map[string]interface{}) ([]models.Category, error) {
	return s.categoryRepo.FindAll(filters)
}

func (s *categoryService) GetCategoryBySlug(slug string) (*models.Category, error) {
	return s.categoryRepo.FindBySlug(slug)
}
