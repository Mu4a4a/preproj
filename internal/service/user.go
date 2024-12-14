package service

import (
	"preproj/internal/repository"
	"preproj/models"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Create(user *models.User) (int64, error) {
	return u.repo.Create(user)
}

func (u *UserService) GetByID(id int64) (*models.User, error) {
	return u.repo.GetByID(id)
}

func (u *UserService) Update(user *models.User) (int64, error) {
	return u.repo.Update(user)
}

func (u *UserService) Delete(id int64) error {
	return u.repo.Delete(id)
}

func (u *UserService) GetAll() ([]*models.User, error) {
	return u.repo.GetAll()
}

// TODO: сложные запросы
/*func (u *UserService) Create(user *models.User) (int64, error) {
	return u.repo.Create(user)
}*/
