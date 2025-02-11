package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"preproj/internal/models"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *models.User) (int64, error) {
	query := `INSERT INTO users (name, email, created_at, updated_at)VALUES ($1, $2, NOW(), NOW()) RETURNING id;`
	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %v", err)
	}
	return user.ID, nil
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	query := `SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1;`
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("user with id %d not found", id)
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	return user, nil
}

func (r *PostgresUserRepository) Update(ctx context.Context, user models.User) (int64, error) {
	query := `UPDATE users SET name = $2, email = $3, updated_at = NOW() WHERE id = $1 RETURNING id;`
	err := r.db.QueryRowContext(ctx, query, user.ID, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to update user: %v", err)
	}
	return user.ID, nil
}

func (r *PostgresUserRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1;`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}

func (r *PostgresUserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	query := `SELECT id, name, email, created_at, updated_at FROM users;`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %v", err)
	}
	defer rows.Close()
	var users []models.User

	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user rows: %v", err)
		}
		users = append(users, user)
	}
	return users, nil
}
