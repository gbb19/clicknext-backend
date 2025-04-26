package dto

import (
	"clicknext-backend/internal/domain"
	"time"
)

type UserResponse struct {
	UserID    uint      `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCreateRequest struct {
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (r *UserCreateRequest) ToUserDomain() *domain.User {
	return &domain.User{
		Username:     r.Username,
		Email:        r.Email,
		PasswordHash: r.Password,
		FirstName:    r.FirstName,
		LastName:     r.LastName,
	}
}

func FromUserDomain(u *domain.User) *UserResponse {
	return &UserResponse{
		UserID:    u.UserID,
		Username:  u.Username,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
