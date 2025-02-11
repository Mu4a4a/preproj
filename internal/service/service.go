package service

import (
	"context"
	"preproj/internal/models"
	"preproj/internal/repository"
)

type User interface {
	Create(ctx context.Context, user *models.User) (int64, error)
	GetByID(ctx context.Context, id int64) (*models.User, error)
	Update(ctx context.Context, user models.User) (int64, error)
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]models.User, error)
}

type Product interface {
	Create(ctx context.Context, product *models.Product) (int64, error)
	GetByID(ctx context.Context, id int64) (*models.Product, error)
	Update(ctx context.Context, product models.Product) (int64, error)
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]models.Product, error)
	GetAllByUserID(ctx context.Context, userID int64) ([]models.Product, error)
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
