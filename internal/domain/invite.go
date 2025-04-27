package domain

import (
	"time"
)

type InviteStatus string

const (
	Pending  InviteStatus = "pending"
	Accepted InviteStatus = "accepted"
	Declined InviteStatus = "declined"
)

type Invite struct {
	InviteID  uint         `gorm:"primaryKey" json:"invite_id"`
	Status    InviteStatus `gorm:"type:invite_status;default:'pending'" json:"status"`
	BoardID   uint         `gorm:"not null" json:"board_id"`
	InviterID uint         `gorm:"not null" json:"inviter_id"`
	InviteeID uint         `gorm:"not null" json:"invitee_id"`
	CreatedAt time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Associations
	Inviter User  `gorm:"foreignKey:InviterID;references:UserID" json:"inviter"`
	Invitee User  `gorm:"foreignKey:InviteeID;references:UserID" json:"invitee"`
	Board   Board `gorm:"foreignKey:BoardID;references:BoardID" json:"board"`
}
