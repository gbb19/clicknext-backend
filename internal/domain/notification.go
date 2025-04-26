package domain

import (
	"time"

	"gorm.io/gorm"
)

type NotifyType string

const (
	TaskType  NotifyType = "task"
	BoardType NotifyType = "board"
)

type Notification struct {
	NotifyID  uint       `gorm:"primaryKey" json:"notify_id"`
	Message   string     `gorm:"not null" json:"message"`
	IsRead    bool       `gorm:"default:false" json:"is_read"`
	Type      NotifyType `gorm:"type:notify_type;not null" json:"type"`
	UserID    uint       `json:"user_id"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Associations
	User User `gorm:"foreignKey:UserID;references:UserID" json:"user"`
}

func (n *Notification) BeforeCreate(db *gorm.DB) error {
	// You can add logic to handle before creating notification, e.g., validation or checks.
	return nil
}
