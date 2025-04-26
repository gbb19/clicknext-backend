package domain

import (
	"time"
)

type Task struct {
	TaskID    uint      `gorm:"primaryKey" json:"task_id"`
	Name      string    `gorm:"not null" json:"name"`
	Position  int       `gorm:"not null" json:"position"`
	DueDate   time.Time `json:"due_date"`
	StartDate time.Time `gorm:"not null" json:"start_date"`
	ColumnID  uint      `json:"column_id"`
	CreatedBy uint      `json:"created_by"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Associations
	CreatedByUser User   `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_user"`
	Column        Column `gorm:"foreignKey:ColumnID;references:ColumnID" json:"column"`
}
