package dto

import (
	"clicknext-backend/internal/domain"
	"time"
)

type InviteResponse struct {
	InviteID  uint      `json:"invite_id"`
	Status    string    `json:"status"`
	InviterID uint      `json:"inviter_id"`
	InviteeID uint      `json:"invitee_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InviteCreateRequest struct {
	InviterID uint `json:"inviter_id" validate:"required"`
	InviteeID uint `json:"invitee_id" validate:"required"`
}

func (r *InviteCreateRequest) ToInviteDomain() *domain.Invite {
	return &domain.Invite{
		InviterID: r.InviterID,
		InviteeID: r.InviteeID,
		Status:    domain.Pending, // Default status is 'pending'
	}
}

func FromInviteDomain(i *domain.Invite) *InviteResponse {
	return &InviteResponse{
		InviteID:  i.InviteID,
		Status:    string(i.Status),
		InviterID: i.InviterID,
		InviteeID: i.InviteeID,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
}
