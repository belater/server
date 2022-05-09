package service

import (
	"server-service/entity"
	"server-service/repository"
)

type UserService interface {
	GetAllUser() ([]entity.User, error)
	GetUserById(id string) (entity.User, error)
	CreateUser(input entity.UserRegister) error
	UpdateUserById(id string, input entity.UserRegister) error
	DeleteUserById(id string) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) GetAllUser() ([]entity.User, error) {
	return s.userRepository.GetAll()
}

func (s *userService) GetUserById(id string) (entity.User, error) {
	return s.userRepository.GetById(id)
}

func (s *userService) CreateUser(input entity.UserRegister) error {
	return s.userRepository.Create(input)
}

func (s *userService) UpdateUserById(id string, input entity.UserRegister) error {
	return s.userRepository.UpdateById(id, input)
}

func (s *userService) DeleteUserById(id string) error {
	return s.userRepository.DeleteById(id)
}
