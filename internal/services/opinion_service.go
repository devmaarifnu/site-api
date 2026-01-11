package services

import (
	"lpmaarifnu-site-api/internal/models"
	"lpmaarifnu-site-api/internal/repositories"
)

type OpinionService interface {
	GetAllOpinions(offset, limit int, filters map[string]interface{}) ([]models.OpinionArticle, int64, error)
	GetOpinionBySlug(slug string) (*models.OpinionArticle, error)
	IncrementViews(id uint) error
}

type opinionService struct {
	opinionRepo repositories.OpinionRepository
}

func NewOpinionService(opinionRepo repositories.OpinionRepository) OpinionService {
	return &opinionService{opinionRepo: opinionRepo}
}

func (s *opinionService) GetAllOpinions(offset, limit int, filters map[string]interface{}) ([]models.OpinionArticle, int64, error) {
	return s.opinionRepo.FindAll(offset, limit, filters)
}

func (s *opinionService) GetOpinionBySlug(slug string) (*models.OpinionArticle, error) {
	return s.opinionRepo.FindBySlug(slug)
}

func (s *opinionService) IncrementViews(id uint) error {
	return s.opinionRepo.IncrementViews(id)
}
