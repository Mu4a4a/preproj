package service

import (
	"context"
	"preproj/internal/models"
	"preproj/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Create(ctx context.Context, user *models.User) (int64, error) {
	return u.repo.Create(ctx, user)
}

func (u *UserService) GetByID(ctx context.Context, id int64) (*models.User, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *UserService) Update(ctx context.Context, user models.User) (int64, error) {
	return u.repo.Update(ctx, user)
}

func (u *UserService) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *UserService) GetAll(ctx context.Context) ([]models.User, error) {
	return u.repo.GetAll(ctx)
}
