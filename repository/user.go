package repository

import (
	"server-service/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]entity.User, error)
	GetById(id string) (entity.User, error)
	Create(input entity.UserRegister) error
	UpdateById(id string, input entity.UserRegister) error
	DeleteById(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]entity.User, error) {
	var users []entity.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepository) GetById(id string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Create(input entity.UserRegister) error {
	if err := r.db.Create(&input).Error; err != nil {
		return err
	}

	return nil
}

// TODO: update soon
func (r *userRepository) UpdateById(id string, input entity.UserRegister) error {
	return nil
}

func (r *userRepository) DeleteById(id string) error {
	if err := r.db.Delete(&entity.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
