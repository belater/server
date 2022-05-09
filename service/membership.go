package service

import (
	"server-service/repository"
	"server-service/entity"
)

type MembershipService interface {
	GetAllMembership() ([]entity.Membership, error)
	GetMembershipById(id string) (entity.Membership, error)
	CreateMembership(input entity.MembershipInput) (error)
	UpdateMembershipById(id string, input entity.MembershipInput) (error)
	DeleteMembershipById(id string) (error)
}

type membershipService struct {
	membershipRepository repository.MembershipRepository
}

func NewMembershipService(membershipRepository repository.MembershipRepository) *membershipService {
	return &membershipService{
		membershipRepository: membershipRepository,
	}
}

func (s *membershipService) GetAllMembership() ([]entity.Membership, error) {
	return s.membershipRepository.GetAll()
}

func (s *membershipService) GetMembershipById(id string) (entity.Membership, error) {
	return s.membershipRepository.GetById(id)
}

func (s *membershipService) CreateMembership(input entity.MembershipInput) (error) {
	return s.membershipRepository.Create(input)
}

func (s *membershipService) UpdateMembershipById(id string, input entity.MembershipInput) (error) {
	return s.membershipRepository.UpdateById(id, input)
}

func (s *membershipService) DeleteMembershipById(id string) (error) {
	return s.membershipRepository.DeleteById(id)
}