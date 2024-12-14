package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"preproj/models"
)

type PostgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {
	return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) Create(product *models.Product) (int64, error) {
	query := `INSERT INTO products (name, description, price, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id;`
	err := r.db.QueryRow(query, product.Name, product.Description, product.Price, product.UserID).Scan(&product.ID)
	if err != nil {
		return 0, fmt.Errorf("error creating product: %v", err)
	}
	return product.ID, nil
}

func (r *PostgresProductRepository) GetByID(id int64) (*models.Product, error) {
	query := `SELECT id, name, description, price, user_id, created_at, updated_at FROM products WHERE id = $1;`
	product := &models.Product{}
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.UserID, &product.CreatedAt, &product.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("product with id %d not found", id)
	} else if err != nil {
		return nil, fmt.Errorf("error getting product: %v", err)
	}
	return product, nil
}

func (r *PostgresProductRepository) Update(product *models.Product) (int64, error) {
	query := `UPDATE products SET name = $1, description = $2, price = $3, user_id = $4, updated_at = NOW() WHERE id = $3 RETURNING id;`
	err := r.db.QueryRow(query, product.Name, product.Description, product.Price, product.UserID).Scan(&product.ID)
	if err != nil {
		return 0, fmt.Errorf("error updating product: %v", err)
	}
	return product.ID, nil
}

func (r *PostgresProductRepository) Delete(id int64) error {
	query := `DELETE FROM products WHERE id = $1;`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting product: %v", err)
	}
	return nil
}

func (r *PostgresProductRepository) GetAll() ([]*models.Product, error) {
	query := `SELECT id, name, description, price, user_id, created_at, updated_at FROM products;`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error getting products: %v", err)
	}
	defer rows.Close()
	var products []*models.Product

	for rows.Next() {
		product := &models.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.UserID, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scaninng user row: %v", err)
		}
		products = append(products, product)
	}
	return products, nil
}
