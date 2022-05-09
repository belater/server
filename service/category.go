package service

import (
	"server-service/repository"
	"server-service/entity"
)

type CategoryService interface {
	GetAllCategory() ([]entity.Category, error)
	GetCategoryById(id string) (entity.Category, error)
	CreateCategory(input entity.CategoryInput) (error)
	UpdateCategoryById(id string, input entity.CategoryInput) (error)
	DeleteCategoryById(id string) (error)
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) *categoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *categoryService) GetAllCategory() ([]entity.Category, error) {
	return s.categoryRepository.GetAll()
}

func (s *categoryService) GetCategoryById(id string) (entity.Category, error) {
	return s.categoryRepository.GetById(id)
}

func (s *categoryService) CreateCategory(input entity.CategoryInput) (error) {
	return s.categoryRepository.Create(input)
}

func (s *categoryService) UpdateCategoryById(id string, input entity.CategoryInput) (error) {
	return s.categoryRepository.UpdateById(id, input)
}

func (s *categoryService) DeleteCategoryById(id string) (error) {
	return s.categoryRepository.DeleteById(id)
}