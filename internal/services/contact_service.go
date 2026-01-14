package services

import (
	"lpmaarifnu-site-api/internal/models"
	"lpmaarifnu-site-api/internal/repositories"
	"time"
)

type ContactService interface {
	SubmitContactMessage(req *models.ContactRequest, ipAddress, userAgent string) (*models.ContactMessage, error)
}

type contactService struct {
	contactRepo repositories.ContactRepository
}

func NewContactService(contactRepo repositories.ContactRepository) ContactService {
	return &contactService{contactRepo: contactRepo}
}

func (s *contactService) SubmitContactMessage(req *models.ContactRequest, ipAddress, userAgent string) (*models.ContactMessage, error) {
	// Generate ticket ID
	ticketID, err := s.contactRepo.GenerateTicketID()
	if err != nil {
		return nil, err
	}

	// Create contact message
	contact := &models.ContactMessage{
		TicketID:  ticketID,
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Subject:   req.Subject,
		Message:   req.Message,
		Status:    "new",
		Priority:  "medium",
		IPAddress: &ipAddress,
		UserAgent: &userAgent,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.contactRepo.Create(contact)
	if err != nil {
		return nil, err
	}

	return contact, nil
}
