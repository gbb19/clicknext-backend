package dto

import (
	"clicknext-backend/internal/domain"
	"time"
)

type BoardMemberResponse struct {
	MemberID  uint      `json:"member_id"`
	UserID    uint      `json:"user_id"`
	BoardID   uint      `json:"board_id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BoardMemberCreateRequest struct {
	UserID  uint   `json:"user_id" validate:"required"`
	BoardID uint   `json:"board_id" validate:"required"`
	Role    string `json:"role" validate:"required,oneof=admin member"`
}

func (r *BoardMemberCreateRequest) ToBoardMemberDomain() *domain.BoardMember {
	role := domain.RoleBoard(r.Role)
	return &domain.BoardMember{
		UserID:  r.UserID,
		BoardID: r.BoardID,
		Role:    role,
	}
}

func FromBoardMemberDomain(bm *domain.BoardMember) *BoardMemberResponse {
	return &BoardMemberResponse{
		MemberID:  bm.MemberID,
		UserID:    bm.UserID,
		BoardID:   bm.BoardID,
		Role:      string(bm.Role),
		CreatedAt: bm.CreatedAt,
		UpdatedAt: bm.UpdatedAt,
	}
}
