package postgres

import (
	"clicknext-backend/internal/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository คือ constructor
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser สร้าง user ใหม่
func (r *UserRepository) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

// GetUserByID หา user จาก id
func (r *UserRepository) GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByUsername หา user จาก username
func (r *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ListUsers ดึง user list ทั้งหมด
func (r *UserRepository) ListUsers() ([]*domain.User, error) {
	var users []*domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser อัปเดต user
func (r *UserRepository) UpdateUser(user *domain.User) error {
	return r.db.Save(user).Error
}

// DeleteUser ลบ user
func (r *UserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}
