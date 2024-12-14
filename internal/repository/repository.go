package repository

import (
	"database/sql"
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

type Repository struct {
	User
	Product
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User:    NewPostgresUserRepository(db),
		Product: NewPostgresProductRepository(db),
	}
}
