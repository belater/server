package repository

import (
	"server-service/entity" 

	"gorm.io/gorm"
)

type ContactRepository interface {
	GetAll() ([]entity.Contact, error)
	GetById(id string) (entity.Contact, error)
	Create(input entity.ContactInput) (error)
	UpdateById(id string, input entity.ContactInput) (error)
	DeleteById(id string) (error)
}

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *contactRepository {
	return &contactRepository{db: db}
}

func (r *contactRepository) GetAll() ([]entity.Contact, error) {
	var contacts []entity.Contact
	
	if err := r.db.Find(&contacts).Error; err != nil {
		return contacts, err
	}

	return contacts, nil
}

func (r *contactRepository) GetById(id string) (entity.Contact, error) {
	var contact entity.Contact
	
	if err := r.db.Where("id = ?", id).Find(&contact).Error; err != nil {
		return contact, err
	}

	return contact, nil
}

func (r *contactRepository) Create(input entity.ContactInput) (error) {
	if err := r.db.Create(&input).Error ; err != nil {
		return err
	}

	return nil
}

// TODO: update soon
func (r *contactRepository) UpdateById(id string, input entity.ContactInput) (error) {
	return nil
}

func (r *contactRepository) DeleteById(id string) (error) {
	if err := r.db.Delete(&entity.Contact{}, id).Error; err != nil {
		return err
	}

	return nil
}