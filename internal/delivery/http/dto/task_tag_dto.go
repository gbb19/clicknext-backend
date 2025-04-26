package dto

import (
	"clicknext-backend/internal/domain"
	"time"
)

type TaskTagResponse struct {
	TaskTagID uint      `json:"task_tag_id"`
	TaskID    uint      `json:"task_id"`
	TagID     uint      `json:"tag_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskTagCreateRequest struct {
	TaskID uint `json:"task_id" validate:"required"`
	TagID  uint `json:"tag_id" validate:"required"`
}

func (r *TaskTagCreateRequest) ToTaskTagDomain() *domain.TaskTag {
	return &domain.TaskTag{
		TaskID: r.TaskID,
		TagID:  r.TagID,
	}
}

func FromTaskTagDomain(t *domain.TaskTag) *TaskTagResponse {
	return &TaskTagResponse{
		TaskTagID: t.TaskTagID,
		TaskID:    t.TaskID,
		TagID:     t.TagID,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
