package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"preproj/internal/models"
)

type PostgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {
	return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) Create(ctx context.Context, product *models.Product) (int64, error) {
	query := `INSERT INTO products (name, description, price, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id;`
	err := r.db.QueryRowContext(ctx, query, product.Name, product.Description, product.Price, product.UserID).Scan(&product.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to create product: %v", err)
	}
	return product.ID, nil
}

func (r *PostgresProductRepository) GetByID(ctx context.Context, id int64) (*models.Product, error) {
	query := `SELECT id, name, description, price, user_id, created_at, updated_at FROM products WHERE id = $1;`
	product := &models.Product{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.UserID, &product.CreatedAt, &product.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("product with id %d not found", id)
	} else if err != nil {
		return nil, fmt.Errorf("failed to get product: %v", err)
	}
	return product, nil
}

func (r *PostgresProductRepository) Update(ctx context.Context, product models.Product) (int64, error) {
	query := `UPDATE products SET name = $2, description = $3, price = $4, user_id = $5, updated_at = NOW() WHERE id = $1 RETURNING id;`
	err := r.db.QueryRowContext(ctx, query, product.ID, product.Name, product.Description, product.Price, product.UserID).Scan(&product.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to update product: %v", err)
	}
	return product.ID, nil
}

func (r *PostgresProductRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM products WHERE id = $1;`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %v", err)
	}
	return nil
}

func (r *PostgresProductRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	query := `SELECT id, name, description, price, user_id, created_at, updated_at FROM products;`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %v", err)
	}
	defer rows.Close()
	var products []models.Product

	for rows.Next() {
		product := models.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.UserID, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *PostgresProductRepository) GetAllByUserID(ctx context.Context, userID int64) ([]models.Product, error) {
	query := `SELECT id, name, description, price, user_id, created_at, updated_at FROM products WHERE user_id = $1`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get products by userID: %v", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		product := models.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.UserID, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		products = append(products, product)
	}
	return products, nil
}
