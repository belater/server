package repository

import (
	"server-service/entity" 

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll() ([]entity.Category, error)
	GetById(id string) (entity.Category, error)
	Create(input entity.CategoryInput) (error)
	UpdateById(id string, input entity.CategoryInput) (error)
	DeleteById(id string) (error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAll() ([]entity.Category, error) {
	var categories []entity.Category
	
	if err := r.db.Find(&categories).Error; err != nil {
		return categories, err
	}

	return categories, nil
}

func (r *categoryRepository) GetById(id string) (entity.Category, error) {
	var category entity.Category
	
	if err := r.db.Where("id = ?", id).Find(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) Create(input entity.CategoryInput) (error) {
	if err := r.db.Create(&input).Error ; err != nil {
		return err
	}

	return nil
}

// TODO: update soon
func (r *categoryRepository) UpdateById(id string, input entity.CategoryInput) (error) {
	return nil
}

func (r *categoryRepository) DeleteById(id string) (error) {
	if err := r.db.Delete(&entity.Category{}, id).Error; err != nil {
		return err
	}

	return nil
}