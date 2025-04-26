package dto

import (
	"clicknext-backend/internal/domain"
	"time"
)

type AssigneeTaskResponse struct {
	AssigneeTaskID uint      `json:"assignee_task_id"`
	AssigneeID     uint      `json:"assignee_id"`
	TaskID         uint      `json:"task_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type AssigneeTaskCreateRequest struct {
	AssigneeID uint `json:"assignee_id" validate:"required"`
	TaskID     uint `json:"task_id" validate:"required"`
}

func (r *AssigneeTaskCreateRequest) ToAssigneeTaskDomain() *domain.AssigneeTask {
	return &domain.AssigneeTask{
		AssigneeID: r.AssigneeID,
		TaskID:     r.TaskID,
	}
}

func FromAssigneeTaskDomain(a *domain.AssigneeTask) *AssigneeTaskResponse {
	return &AssigneeTaskResponse{
		AssigneeTaskID: a.AssigneeTaskID,
		AssigneeID:     a.AssigneeID,
		TaskID:         a.TaskID,
		CreatedAt:      a.CreatedAt,
		UpdatedAt:      a.UpdatedAt,
	}
}
