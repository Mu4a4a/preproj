package service

import (
	"context"
	"preproj/internal/models"
	"preproj/internal/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (u *ProductService) Create(ctx context.Context, product *models.Product) (int64, error) {
	return u.repo.Create(ctx, product)
}

func (u *ProductService) GetByID(ctx context.Context, id int64) (*models.Product, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *ProductService) Update(ctx context.Context, product models.Product) (int64, error) {
	return u.repo.Update(ctx, product)
}

func (u *ProductService) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *ProductService) GetAll(ctx context.Context) ([]models.Product, error) {
	return u.repo.GetAll(ctx)
}

func (u *ProductService) GetAllByUserID(ctx context.Context, userID int64) ([]models.Product, error) {
	return u.repo.GetAllByUserID(ctx, userID)
}
