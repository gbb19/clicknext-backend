package usecase

import (
	"clicknext-backend/internal/domain"
)

type TaskRepository interface {
	CreateTask(task *domain.Task) error
	GetTaskByID(id uint) (*domain.Task, error)
	GetTasksByColumnID(columnID uint) ([]*domain.Task, error)
	UpdateTask(task *domain.Task) error
	DeleteTask(id uint) error
	UpdateTaskPosition(taskID uint, newPosition int, columnID uint) error
}

type TaskUsecase struct {
	repo TaskRepository
}

func NewTaskUsecase(repo TaskRepository) *TaskUsecase {
	return &TaskUsecase{repo: repo}
}

func (uc *TaskUsecase) CreateTask(task *domain.Task) error {
	return uc.repo.CreateTask(task)
}

func (uc *TaskUsecase) GetTaskByID(id uint) (*domain.Task, error) {
	return uc.repo.GetTaskByID(id)
}

func (uc *TaskUsecase) GetTasksByColumnID(columnID uint) ([]*domain.Task, error) {
	return uc.repo.GetTasksByColumnID(columnID)
}

func (uc *TaskUsecase) UpdateTask(task *domain.Task) error {
	return uc.repo.UpdateTask(task)
}

func (uc *TaskUsecase) DeleteTask(id uint) error {
	return uc.repo.DeleteTask(id)
}

func (uc *TaskUsecase) UpdateTaskPosition(taskID uint, newPosition int, columnID uint) error {
	return uc.repo.UpdateTaskPosition(taskID, newPosition, columnID)
}
