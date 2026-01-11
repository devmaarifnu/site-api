package services

import (
	"lpmaarifnu-site-api/internal/models"
	"lpmaarifnu-site-api/internal/repositories"
)

type PageService interface {
	GetPageBySlug(slug string) (*models.Page, error)
}

type pageService struct {
	pageRepo repositories.PageRepository
}

func NewPageService(pageRepo repositories.PageRepository) PageService {
	return &pageService{pageRepo: pageRepo}
}

func (s *pageService) GetPageBySlug(slug string) (*models.Page, error) {
	return s.pageRepo.FindBySlug(slug)
}
