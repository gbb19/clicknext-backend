package domain

import (
	"time"
)

type TaskTag struct {
	TaskTagID uint      `gorm:"primaryKey" json:"task_tag_id"`
	TaskID    uint      `json:"task_id"`
	TagID     uint      `json:"tag_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Associations
	Task Task `gorm:"foreignKey:TaskID;references:TaskID" json:"task"`
	Tag  Tag  `gorm:"foreignKey:TagID;references:TagID" json:"tag"`
}
