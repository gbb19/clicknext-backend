package usecase

import (
	"clicknext-backend/internal/domain"
	"clicknext-backend/internal/repository/postgres"
	"clicknext-backend/internal/utils"
	"errors"

	"gorm.io/gorm"
)

type AuthUseCase struct {
	userRepo *postgres.UserRepository
}

func NewAuthUseCase(userRepo *postgres.UserRepository) *AuthUseCase {
	return &AuthUseCase{userRepo: userRepo}
}

// Register สำหรับลงทะเบียนผู้ใช้ใหม่
func (uc *AuthUseCase) Register(user *domain.User) error {
	return uc.userRepo.CreateUser(user)

}

// Login สำหรับเข้าสู่ระบบ
func (uc *AuthUseCase) Login(username string, password string) (*domain.User, string, error) {
	user, err := uc.userRepo.GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", errors.New("invalid username or password")
		}
		return nil, "", err
	}

	if !user.CheckPassword(password) {
		return nil, "", errors.New("invalid username or password")
	}

	token, err := utils.GenerateAccessToken(user.UserID, user.Username, user.FirstName, user.LastName)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
