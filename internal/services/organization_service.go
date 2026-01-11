package services

import (
	"lpmaarifnu-site-api/internal/models"
	"lpmaarifnu-site-api/internal/repositories"
)

type OrganizationService interface {
	GetOrganizationStructure() (map[string]interface{}, error)
	GetBoardMembers(filters map[string]interface{}) ([]models.BoardMember, error)
}

type organizationService struct {
	orgRepo repositories.OrganizationRepository
}

func NewOrganizationService(orgRepo repositories.OrganizationRepository) OrganizationService {
	return &organizationService{orgRepo: orgRepo}
}

func (s *organizationService) GetOrganizationStructure() (map[string]interface{}, error) {
	return s.orgRepo.GetOrganizationStructure()
}

func (s *organizationService) GetBoardMembers(filters map[string]interface{}) ([]models.BoardMember, error) {
	return s.orgRepo.FindBoardMembers(filters)
}
