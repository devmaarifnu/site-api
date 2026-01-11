package services

import (
	"lpmaarifnu-site-api/internal/models"
	"lpmaarifnu-site-api/internal/repositories"
)

type DocumentService interface {
	GetAllDocuments(offset, limit int, filters map[string]interface{}) ([]models.Document, int64, error)
	GetDocumentByID(id uint) (*models.Document, error)
	IncrementDownloads(id uint) error
}

type documentService struct {
	docRepo repositories.DocumentRepository
}

func NewDocumentService(docRepo repositories.DocumentRepository) DocumentService {
	return &documentService{docRepo: docRepo}
}

func (s *documentService) GetAllDocuments(offset, limit int, filters map[string]interface{}) ([]models.Document, int64, error) {
	return s.docRepo.FindAll(offset, limit, filters)
}

func (s *documentService) GetDocumentByID(id uint) (*models.Document, error) {
	return s.docRepo.FindByID(id)
}

func (s *documentService) IncrementDownloads(id uint) error {
	return s.docRepo.IncrementDownloads(id)
}
