package dto

import (
	"clicknext-backend/internal/domain"
	"time"
)

type TagResponse struct {
	TagID     uint      `json:"tag_id"`
	Name      string    `json:"name"`
	CreatedBy uint      `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

type TagCreateRequest struct {
	Name      string `json:"name" validate:"required"`
	CreatedBy uint   `json:"created_by" validate:"required"`
}

func (r *TagCreateRequest) ToTagDomain() *domain.Tag {
	return &domain.Tag{
		Name:      r.Name,
		CreatedBy: r.CreatedBy,
	}
}

func FromTagDomain(t *domain.Tag) *TagResponse {
	return &TagResponse{
		TagID:     t.TagID,
		Name:      t.Name,
		CreatedBy: t.CreatedBy,
		CreatedAt: t.CreatedAt,
	}
}
