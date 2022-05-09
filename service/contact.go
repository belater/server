package service

import (
	"server-service/repository"
	"server-service/entity"
)

type ContactService interface {
	GetAllContact() ([]entity.Contact, error)
	GetContactById(id string) (entity.Contact, error)
	CreateContact(input entity.ContactInput) (error)
	UpdateContactById(id string, input entity.ContactInput) (error)
	DeleteContactById(id string) (error)
}

type contactService struct {
	contactRepository repository.ContactRepository
}

func NewContactService(contactRepository repository.ContactRepository) *contactService {
	return &contactService{
		contactRepository: contactRepository,
	}
}

func (s *contactService) GetAllContact() ([]entity.Contact, error) {
	return s.contactRepository.GetAll()
}

func (s *contactService) GetContactById(id string) (entity.Contact, error) {
	return s.contactRepository.GetById(id)
}

func (s *contactService) CreateContact(input entity.ContactInput) (error) {
	return s.contactRepository.Create(input)
}

func (s *contactService) UpdateContactById(id string, input entity.ContactInput) (error) {
	return s.contactRepository.UpdateById(id, input)
}

func (s *contactService) DeleteContactById(id string) (error) {
	return s.contactRepository.DeleteById(id)
}