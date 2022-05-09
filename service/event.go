package service

import (
	"server-service/repository"
	"server-service/entity"
)

type EventService interface {
	GetAllEvent() ([]entity.Event, error)
	GetEventById(id string) (entity.Event, error)
	CreateEvent(input entity.EventInput) (error)
	UpdateEventById(id string, input entity.EventInput) (error)
	DeleteEventById(id string) (error)
}

type eventService struct {
	eventRepository repository.EventRepository
}

func NewEventService(eventRepository repository.EventRepository) *eventService {
	return &eventService{
		eventRepository: eventRepository,
	}
}

func (s *eventService) GetAllEvent() ([]entity.Event, error) {
	return s.eventRepository.GetAll()
}

func (s *eventService) GetEventById(id string) (entity.Event, error) {
	return s.eventRepository.GetById(id)
}

func (s *eventService) CreateEvent(input entity.EventInput) (error) {
	return s.eventRepository.Create(input)
}

func (s *eventService) UpdateEventById(id string, input entity.EventInput) (error) {
	return s.eventRepository.UpdateById(id, input)
}

func (s *eventService) DeleteEventById(id string) (error) {
	return s.eventRepository.DeleteById(id)
}