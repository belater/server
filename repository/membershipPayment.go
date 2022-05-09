package repository

import (
	"server-service/entity" 

	"gorm.io/gorm"
)

type MembershipPaymentRepository interface {
	GetAll() ([]entity.MembershipPayment, error)
	GetById(id string) (entity.MembershipPayment, error)
	Create(input entity.MembershipPaymentInput) (error)
	UpdateById(id string, input entity.MembershipPaymentInput) (error)
	DeleteById(id string) (error)
}

type membershipPaymentRepository struct {
	db *gorm.DB
}

func NewMembershipPaymentRepository(db *gorm.DB) *membershipPaymentRepository {
	return &membershipPaymentRepository{db: db}
}

func (r *membershipPaymentRepository) GetAll() ([]entity.MembershipPayment, error) {
	var membershipPayments []entity.MembershipPayment
	
	if err := r.db.Find(&membershipPayments).Error; err != nil {
		return membershipPayments, err
	}

	return membershipPayments, nil
}

func (r *membershipPaymentRepository) GetById(id string) (entity.MembershipPayment, error) {
	var membershipPayment entity.MembershipPayment
	
	if err := r.db.Where("id = ?", id).Find(&membershipPayment).Error; err != nil {
		return membershipPayment, err
	}

	return membershipPayment, nil
}

func (r *membershipPaymentRepository) Create(input entity.MembershipPaymentInput) (error) {
	if err := r.db.Create(&input).Error ; err != nil {
		return err
	}

	return nil
}

// TODO: update soon
func (r *membershipPaymentRepository) UpdateById(id string, input entity.MembershipPaymentInput) (error) {
	return nil
}

func (r *membershipPaymentRepository) DeleteById(id string) (error) {
	if err := r.db.Delete(&entity.MembershipPayment{}, id).Error; err != nil {
		return err
	}

	return nil
}