package domain

import (
	"time"
)

type RoleBoard string

const (
	Admin  RoleBoard = "admin"
	Member RoleBoard = "member"
)

type BoardMember struct {
	MemberID  uint      `gorm:"primaryKey" json:"member_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	BoardID   uint      `gorm:"not null" json:"board_id"`
	Role      RoleBoard `gorm:"type:role_board;default:'member'" json:"role"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Associations
	User  User  `gorm:"foreignKey:UserID;references:UserID" json:"user"`
	Board Board `gorm:"foreignKey:BoardID;references:BoardID" json:"board"`
}
