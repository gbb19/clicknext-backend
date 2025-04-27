package domain

import (
	"time"
)

type Tag struct {
	TagID     uint      `gorm:"primaryKey" json:"tag_id"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedBy uint      `json:"created_by"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`

	// Associations
	User User `gorm:"foreignKey:CreatedBy;references:UserID"`
}
