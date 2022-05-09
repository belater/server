package repository

import (
	"server-service/entity" 

	"gorm.io/gorm"
)

type EventRepository interface {
	GetAll() ([]entity.Event, error)
	GetById(id string) (entity.Event, error)
	Create(input entity.EventInput) (error)
	UpdateById(id string, input entity.EventInput) (error)
	DeleteById(id string) (error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *eventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) GetAll() ([]entity.Event, error) {
	var events []entity.Event
	
	if err := r.db.Find(&events).Error; err != nil {
		return events, err
	}

	return events, nil
}

func (r *eventRepository) GetById(id string) (entity.Event, error) {
	var event entity.Event
	
	if err := r.db.Where("id = ?", id).Find(&event).Error; err != nil {
		return event, err
	}

	return event, nil
}

func (r *eventRepository) Create(input entity.EventInput) (error) {
	if err := r.db.Create(&input).Error ; err != nil {
		return err
	}

	return nil
}

// TODO: update soon
func (r *eventRepository) UpdateById(id string, input entity.EventInput) (error) {
	return nil
}

func (r *eventRepository) DeleteById(id string) (error) {
	if err := r.db.Delete(&entity.Event{}, id).Error; err != nil {
		return err
	}

	return nil
}