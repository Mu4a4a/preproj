package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"preproj/models"
)

// TODO: -warping errors, обработка ошибок, сделать больше кейсов с проверкой ошибки, проверка входных данных

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(user *models.User) (int64, error) {
	query := `INSERT INTO users (name, email, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING id;`
	err := r.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return 0, fmt.Errorf("error creating user: %v", err)
	}
	return user.ID, nil
}

func (r *PostgresUserRepository) GetByID(id int64) (*models.User, error) {
	query := `SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1;`
	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("user with id %d not found", id)
	} else if err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}
	return user, nil
}

func (r *PostgresUserRepository) Update(user *models.User) (int64, error) {
	query := `UPDATE users SET name = $1, email = $2, updated_at = NOW() WHERE id = $3 RETURNING id;`
	err := r.db.QueryRow(query, user.ID, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return 0, fmt.Errorf("error updating user: %v", err)
	}
	return user.ID, nil
}

func (r *PostgresUserRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = $1;`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	return nil
}

func (r *PostgresUserRepository) GetAll() ([]*models.User, error) {
	query := `SELECT id, name, email, created_at, updated_at FROM users;`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error getting users: %v", err)
	}
	defer rows.Close()
	var users []*models.User

	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scaninng user row: %v", err)
		}
		users = append(users, user)
	}
	return users, nil
}
