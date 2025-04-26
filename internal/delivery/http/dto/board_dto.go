package dto

import (
	"clicknext-backend/internal/domain"
	"time"
)

type BoardResponse struct {
	BoardID     uint      `json:"board_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedBy   uint      `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BoardCreateRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	CreatedBy   uint   `json:"created_by" validate:"required"`
}

func (r *BoardCreateRequest) ToBoardDomain() *domain.Board {
	return &domain.Board{
		Title:       r.Title,
		Description: r.Description,
		CreatedBy:   r.CreatedBy,
	}
}

func FromBoardDomain(b *domain.Board) *BoardResponse {
	return &BoardResponse{
		BoardID:     b.BoardID,
		Title:       b.Title,
		Description: b.Description,
		CreatedBy:   b.CreatedBy,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}
