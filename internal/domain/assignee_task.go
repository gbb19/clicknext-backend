package domain

import (
	"time"
)

type AssigneeTask struct {
	AssigneeTaskID uint      `gorm:"primaryKey" json:"assignee_task_id"`
	AssigneeID     uint      `json:"assignee_id"`
	TaskID         uint      `json:"task_id"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Associations
	User User `gorm:"foreignKey:AssigneeID;references:UserID" json:"user"`
	Task Task `gorm:"foreignKey:TaskID;references:TaskID" json:"task"`
}
