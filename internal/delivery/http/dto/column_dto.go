package dto

import (
	"clicknext-backend/internal/domain"
	"time"
)

type ColumnResponse struct {
	ColumnID  uint      `json:"column_id"`
	Name      string    `json:"name"`
	Position  int       `json:"position"`
	Color     string    `json:"color"`
	CreatedBy uint      `json:"created_by"`
	BoardID   uint      `json:"board_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ColumnCreateRequest struct {
	Name      string `json:"name" validate:"required"`
	Color     string `json:"color" validate:"required"`
	CreatedBy uint   `json:"created_by" validate:"required"`
	BoardID   uint   `json:"board_id" validate:"required"`
}

func (r *ColumnCreateRequest) ToColumnDomain() *domain.Column {
	return &domain.Column{
		Name:      r.Name,
		Color:     r.Color,
		CreatedBy: r.CreatedBy,
		BoardID:   r.BoardID,
	}
}

func FromColumnDomain(c *domain.Column) *ColumnResponse {
	return &ColumnResponse{
		ColumnID:  c.ColumnID,
		Name:      c.Name,
		Color:     c.Color,
		Position:  c.Position,
		CreatedBy: c.CreatedBy,
		BoardID:   c.BoardID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
