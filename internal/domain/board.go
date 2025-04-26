package domain

import (
	"time"
)

type Board struct {
	BoardID     uint      `gorm:"primaryKey" json:"board_id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"size:255" json:"description"`
	CreatedBy   uint      `gorm:"not null" json:"created_by"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Association with User (created_by)
	User User `gorm:"foreignKey:CreatedBy;references:UserID"`
}
