package services

import (
	"lpmaarifnu-site-api/internal/models"
	"lpmaarifnu-site-api/internal/repositories"
)

type PengurusService interface {
	GetAllPengurus(filters map[string]interface{}) ([]models.Pengurus, error)
}

type pengurusService struct {
	pengurusRepo repositories.PengurusRepository
}

func NewPengurusService(pengurusRepo repositories.PengurusRepository) PengurusService {
	return &pengurusService{pengurusRepo: pengurusRepo}
}

func (s *pengurusService) GetAllPengurus(filters map[string]interface{}) ([]models.Pengurus, error) {
	return s.pengurusRepo.FindAll(filters)
}
