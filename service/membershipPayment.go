package service

import (
	"server-service/repository"
	"server-service/entity"
)

type MembershipPaymentService interface {
	GetAllMembershipPayment() ([]entity.MembershipPayment, error)
	GetMembershipPaymentById(id string) (entity.MembershipPayment, error)
	CreateMembershipPayment(input entity.MembershipPaymentInput) (error)
	UpdateMembershipPaymentById(id string, input entity.MembershipPaymentInput) (error)
	DeleteMembershipPaymentById(id string) (error)
}

type membershipPaymentService struct {
	membershipPaymentRepository repository.MembershipPaymentRepository
}

func NewMembershipPaymentService(membershipPaymentRepository repository.MembershipPaymentRepository) *membershipPaymentService {
	return &membershipPaymentService{
		membershipPaymentRepository: membershipPaymentRepository,
	}
}

func (s *membershipPaymentService) GetAllMembershipPayment() ([]entity.MembershipPayment, error) {
	return s.membershipPaymentRepository.GetAll()
}

func (s *membershipPaymentService) GetMembershipPaymentById(id string) (entity.MembershipPayment, error) {
	return s.membershipPaymentRepository.GetById(id)
}

func (s *membershipPaymentService) CreateMembershipPayment(input entity.MembershipPaymentInput) (error) {
	return s.membershipPaymentRepository.Create(input)
}

func (s *membershipPaymentService) UpdateMembershipPaymentById(id string, input entity.MembershipPaymentInput) (error) {
	return s.membershipPaymentRepository.UpdateById(id, input)
}

func (s *membershipPaymentService) DeleteMembershipPaymentById(id string) (error) {
	return s.membershipPaymentRepository.DeleteById(id)
}