package service

import (
	"preproj/internal/repository"
	"preproj/models"
)

type User interface {
	Create(user *models.User) (int64, error)
	GetByID(id int64) (*models.User, error)
	Update(user *models.User) (int64, error)
	Delete(id int64) error
	GetAll() ([]*models.User, error)
}

type Product interface {
	Create(product *models.Product) (int64, error)
	GetByID(id int64) (*models.Product, error)
	Update(product *models.Product) (int64, error)
	Delete(id int64) error
	GetAll() ([]*models.Product, error)
}

type Service struct {
	User
	Product
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:    NewUserService(repo.User),
		Product: NewProductService(repo.Product),
	}
}
