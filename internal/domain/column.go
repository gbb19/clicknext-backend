package domain

import (
	"time"
)

type Column struct {
	ColumnID  uint      `gorm:"primaryKey" json:"column_id"`
	Name      string    `gorm:"not null" json:"name"`
	Color     string    `gorm:"default:'#5D5D5D';not null" json:"color"`
	Position  int       `gorm:"not null" json:"position"`
	CreatedBy uint      `json:"created_by"`
	BoardID   uint      `json:"board_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Associations
	CreatedByUser User  `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_user"`
	Board         Board `gorm:"foreignKey:BoardID;references:BoardID" json:"board"`
}
