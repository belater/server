package repository

import (
	"server-service/entity" 

	"gorm.io/gorm"
)

type MembershipRepository interface {
	GetAll() ([]entity.Membership, error)
	GetById(id string) (entity.Membership, error)
	Create(input entity.MembershipInput) (error)
	UpdateById(id string, input entity.MembershipInput) (error)
	DeleteById(id string) (error)
}

type membershipRepository struct {
	db *gorm.DB
}

func NewMembershipRepository(db *gorm.DB) *membershipRepository {
	return &membershipRepository{db: db}
}

func (r *membershipRepository) GetAll() ([]entity.Membership, error) {
	var memberships []entity.Membership
	
	if err := r.db.Find(&memberships).Error; err != nil {
		return memberships, err
	}

	return memberships, nil
}

func (r *membershipRepository) GetById(id string) (entity.Membership, error) {
	var membership entity.Membership
	
	if err := r.db.Where("id = ?", id).Find(&membership).Error; err != nil {
		return membership, err
	}

	return membership, nil
}

func (r *membershipRepository) Create(input entity.MembershipInput) (error) {
	if err := r.db.Create(&input).Error ; err != nil {
		return err
	}

	return nil
}

// TODO: update soon
func (r *membershipRepository) UpdateById(id string, input entity.MembershipInput) (error) {
	return nil
}

func (r *membershipRepository) DeleteById(id string) (error) {
	if err := r.db.Delete(&entity.Membership{}, id).Error; err != nil {
		return err
	}

	return nil
}