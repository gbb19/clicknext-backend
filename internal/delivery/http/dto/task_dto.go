package dto

import (
	"clicknext-backend/internal/domain"
	"time"
)

type TaskResponse struct {
	TaskID    uint      `json:"task_id"`
	Name      string    `json:"name"`
	Position  int       `json:"position"`
	DueDate   time.Time `json:"due_date"`
	StartDate time.Time `json:"start_date"`
	ColumnID  uint      `json:"column_id"`
	CreatedBy uint      `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskCreateRequest struct {
	Name      string    `json:"name" validate:"required"`
	Position  int       `json:"position" validate:"required"`
	DueDate   time.Time `json:"due_date"`
	StartDate time.Time `json:"start_date" validate:"required"`
	ColumnID  uint      `json:"column_id" validate:"required"`
	CreatedBy uint      `json:"created_by" validate:"required"`
}

func (r *TaskCreateRequest) ToTaskDomain() *domain.Task {
	return &domain.Task{
		Name:      r.Name,
		Position:  r.Position,
		DueDate:   r.DueDate,
		StartDate: r.StartDate,
		ColumnID:  r.ColumnID,
		CreatedBy: r.CreatedBy,
	}
}

func FromTaskDomain(t *domain.Task) *TaskResponse {
	return &TaskResponse{
		TaskID:    t.TaskID,
		Name:      t.Name,
		Position:  t.Position,
		DueDate:   t.DueDate,
		StartDate: t.StartDate,
		ColumnID:  t.ColumnID,
		CreatedBy: t.CreatedBy,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
