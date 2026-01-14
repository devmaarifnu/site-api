package repositories

import (
	"fmt"
	"lpmaarifnu-site-api/internal/models"

	"gorm.io/gorm"
)

type ContactRepository interface {
	Create(contact *models.ContactMessage) error
	GenerateTicketID() (string, error)
}

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{db: db}
}

func (r *contactRepository) Create(contact *models.ContactMessage) error {
	return r.db.Create(contact).Error
}

func (r *contactRepository) GenerateTicketID() (string, error) {
	var count int64
	year := "2024" // You might want to make this dynamic

	err := r.db.Model(&models.ContactMessage{}).
		Where("ticket_id LIKE ?", "CTK-"+year+"-%").
		Count(&count).Error

	if err != nil {
		return "", err
	}

	// Format: CTK-2024-0001
	ticketID := fmt.Sprintf("CTK-%s-%04d", year, count+1)
	return ticketID, nil
}
