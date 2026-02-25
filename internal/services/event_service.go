package services

import (
	"lpmaarifnu-site-api/internal/models"
	"lpmaarifnu-site-api/internal/repositories"
)

type EventService interface {
	GetAllEvents(filters map[string]interface{}) ([]models.EventFlayer, error)
}

type eventService struct {
	eventRepo repositories.EventRepository
}

func NewEventService(eventRepo repositories.EventRepository) EventService {
	return &eventService{eventRepo: eventRepo}
}

func (s *eventService) GetAllEvents(filters map[string]interface{}) ([]models.EventFlayer, error) {
	return s.eventRepo.FindAll(filters)
}
