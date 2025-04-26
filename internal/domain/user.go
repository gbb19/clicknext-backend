package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID       uint      `gorm:"primaryKey" json:"user_id"`
	Username     string    `gorm:"size:100;unique;not null" json:"username"`
	Email        string    `gorm:"size:255;unique;not null" json:"email"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"`
	FirstName    string    `gorm:"size:255" json:"fisrt_name"`
	LastName     string    `gorm:"size:255" json:"last_name"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	validationErrors := make(map[string]string)

	var existingUser User
	if err := db.Where("username = ?", u.Username).First(&existingUser).Error; err == nil {
		validationErrors["Username"] = "username already exists"
	}

	if err := db.Where("email = ?", u.Email).First(&existingUser).Error; err == nil {
		validationErrors["Email"] = "email already exists"
	}

	if len(validationErrors) > 0 {
		return &ValidationError{
			Message: "Validation failed",
			Errors:  validationErrors,
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)
	return nil
}
