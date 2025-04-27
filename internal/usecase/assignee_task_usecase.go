package usecase

import (
	"clicknext-backend/internal/domain"
)

type AssigneeTaskRepository interface {
	AssignTask(assigneeTask *domain.AssigneeTask) error
	UnassignTask(assigneeID, taskID uint) error
	GetAssigneesByTaskID(taskID uint) ([]*domain.AssigneeTask, error)
	GetTasksByAssigneeID(assigneeID uint) ([]*domain.AssigneeTask, error)
}

type AssigneeTaskUsecase struct {
	repo AssigneeTaskRepository
}

func NewAssigneeTaskUsecase(repo AssigneeTaskRepository) *AssigneeTaskUsecase {
	return &AssigneeTaskUsecase{repo: repo}
}

func (uc *AssigneeTaskUsecase) AssignTask(assigneeTask *domain.AssigneeTask) error {
	return uc.repo.AssignTask(assigneeTask)
}

func (uc *AssigneeTaskUsecase) UnassignTask(assigneeID, taskID uint) error {
	return uc.repo.UnassignTask(assigneeID, taskID)
}

func (uc *AssigneeTaskUsecase) GetAssigneesByTaskID(taskID uint) ([]*domain.AssigneeTask, error) {
	return uc.repo.GetAssigneesByTaskID(taskID)
}

func (uc *AssigneeTaskUsecase) GetTasksByAssigneeID(assigneeID uint) ([]*domain.AssigneeTask, error) {
	return uc.repo.GetTasksByAssigneeID(assigneeID)
}
