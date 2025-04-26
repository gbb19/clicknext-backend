package usecase

import (
	"clicknext-backend/internal/domain"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByID(id uint) (*domain.User, error)
}

type UserUsecase struct {
	repo UserRepository
}

func NewUserUsecase(repo UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) CreateUser(user *domain.User) error {
	return uc.repo.CreateUser(user)
}

func (uc *UserUsecase) GetUserByID(id uint) (*domain.User, error) {
	return uc.repo.GetUserByID(id)
}
