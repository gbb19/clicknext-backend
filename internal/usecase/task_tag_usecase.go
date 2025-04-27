package usecase

import (
	"clicknext-backend/internal/domain"
	"errors"
)

type TaskTagRepository interface {
	CreateTaskTag(taskTag *domain.TaskTag) error
	GetTaskTagByID(id uint) (*domain.TaskTag, error)
	GetTaskTagsByTaskID(taskID uint) ([]*domain.TaskTag, error)
	GetTaskTagsByTagID(tagID uint) ([]*domain.TaskTag, error)
	DeleteTaskTag(id uint) error
	DeleteTaskTagByTaskAndTag(taskID, tagID uint) error
	CheckTaskTagExists(taskID uint, tagID uint) (bool, error)
	DeleteTaskTagsByTaskID(taskID uint) error
}

type TaskTagUsecase struct {
	repo     TaskTagRepository
	taskRepo TaskRepository
	tagRepo  TagRepository
}

func NewTaskTagUsecase(repo TaskTagRepository, taskRepo TaskRepository, tagRepo TagRepository) *TaskTagUsecase {
	return &TaskTagUsecase{
		repo:     repo,
		taskRepo: taskRepo,
		tagRepo:  tagRepo,
	}
}

func (uc *TaskTagUsecase) CreateTaskTag(taskTag *domain.TaskTag) error {
	// Check if task exists
	_, err := uc.taskRepo.GetTaskByID(taskTag.TaskID)
	if err != nil {
		return errors.New("task not found")
	}

	// Check if tag exists
	_, err = uc.tagRepo.GetTagByID(taskTag.TagID)
	if err != nil {
		return errors.New("tag not found")
	}

	// Check if task-tag combination already exists
	exists, err := uc.repo.CheckTaskTagExists(taskTag.TaskID, taskTag.TagID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("tag is already assigned to this task")
	}

	return uc.repo.CreateTaskTag(taskTag)
}

func (uc *TaskTagUsecase) GetTaskTagByID(id uint) (*domain.TaskTag, error) {
	return uc.repo.GetTaskTagByID(id)
}

func (uc *TaskTagUsecase) GetTaskTagsByTaskID(taskID uint) ([]*domain.TaskTag, error) {
	// Verify task exists
	_, err := uc.taskRepo.GetTaskByID(taskID)
	if err != nil {
		return nil, errors.New("task not found")
	}

	return uc.repo.GetTaskTagsByTaskID(taskID)
}

func (uc *TaskTagUsecase) GetTaskTagsByTagID(tagID uint) ([]*domain.TaskTag, error) {
	// Verify tag exists
	_, err := uc.tagRepo.GetTagByID(tagID)
	if err != nil {
		return nil, errors.New("tag not found")
	}

	return uc.repo.GetTaskTagsByTagID(tagID)
}

func (uc *TaskTagUsecase) DeleteTaskTag(id uint) error {
	// Check if task tag exists
	_, err := uc.repo.GetTaskTagByID(id)
	if err != nil {
		return errors.New("task tag not found")
	}

	return uc.repo.DeleteTaskTag(id)
}

func (uc *TaskTagUsecase) DeleteTaskTagByTaskAndTag(taskID, tagID uint) error {
	// Verify task exists
	_, err := uc.taskRepo.GetTaskByID(taskID)
	if err != nil {
		return errors.New("task not found")
	}

	// Verify tag exists
	_, err = uc.tagRepo.GetTagByID(tagID)
	if err != nil {
		return errors.New("tag not found")
	}

	return uc.repo.DeleteTaskTagByTaskAndTag(taskID, tagID)
}

func (uc *TaskTagUsecase) DeleteTaskTagsByTaskID(taskID uint) error {
	// Verify task exists
	_, err := uc.taskRepo.GetTaskByID(taskID)
	if err != nil {
		return errors.New("task not found")
	}

	return uc.repo.DeleteTaskTagsByTaskID(taskID)
}
