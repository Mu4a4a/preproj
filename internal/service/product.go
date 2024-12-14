package service

import (
	"preproj/internal/repository"
	"preproj/models"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (u *ProductService) Create(product *models.Product) (int64, error) {
	return u.repo.Create(product)
}

func (u *ProductService) GetByID(id int64) (*models.Product, error) {
	return u.repo.GetByID(id)
}

func (u *ProductService) Update(product *models.Product) (int64, error) {
	return u.repo.Update(product)
}

func (u *ProductService) Delete(id int64) error {
	return u.repo.Delete(id)
}

func (u *ProductService) GetAll() ([]*models.Product, error) {
	return u.repo.GetAll()
}
