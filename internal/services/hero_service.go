package services

import (
	"lpmaarifnu-site-api/internal/models"
	"lpmaarifnu-site-api/internal/repositories"
)

type HeroService interface {
	GetActiveSlides() ([]models.HeroSlide, error)
}

type heroService struct {
	heroRepo repositories.HeroRepository
}

func NewHeroService(heroRepo repositories.HeroRepository) HeroService {
	return &heroService{heroRepo: heroRepo}
}

func (s *heroService) GetActiveSlides() ([]models.HeroSlide, error) {
	return s.heroRepo.FindActiveSlides()
}
