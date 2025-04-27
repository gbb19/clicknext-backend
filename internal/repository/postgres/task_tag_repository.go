package postgres

import (
	"clicknext-backend/internal/domain"

	"gorm.io/gorm"
)

type TaskTagRepository struct {
	db *gorm.DB
}

func NewTaskTagRepository(db *gorm.DB) *TaskTagRepository {
	return &TaskTagRepository{db: db}
}

func (r *TaskTagRepository) CreateTaskTag(taskTag *domain.TaskTag) error {
	return r.db.Create(taskTag).Error
}

func (r *TaskTagRepository) GetTaskTagByID(id uint) (*domain.TaskTag, error) {
	var taskTag domain.TaskTag
	if err := r.db.Preload("Task").Preload("Tag").First(&taskTag, id).Error; err != nil {
		return nil, err
	}
	return &taskTag, nil
}

func (r *TaskTagRepository) GetTaskTagsByTaskID(taskID uint) ([]*domain.TaskTag, error) {
	var taskTags []*domain.TaskTag
	if err := r.db.Where("task_id = ?", taskID).Preload("Tag").Find(&taskTags).Error; err != nil {
		return nil, err
	}
	return taskTags, nil
}

func (r *TaskTagRepository) GetTaskTagsByTagID(tagID uint) ([]*domain.TaskTag, error) {
	var taskTags []*domain.TaskTag
	if err := r.db.Where("tag_id = ?", tagID).Preload("Task").Find(&taskTags).Error; err != nil {
		return nil, err
	}
	return taskTags, nil
}

func (r *TaskTagRepository) DeleteTaskTag(id uint) error {
	return r.db.Delete(&domain.TaskTag{}, id).Error
}

func (r *TaskTagRepository) DeleteTaskTagByTaskAndTag(taskID, tagID uint) error {
	return r.db.Where("task_id = ? AND tag_id = ?", taskID, tagID).Delete(&domain.TaskTag{}).Error
}

func (r *TaskTagRepository) CheckTaskTagExists(taskID uint, tagID uint) (bool, error) {
	var count int64
	err := r.db.Model(&domain.TaskTag{}).Where("task_id = ? AND tag_id = ?", taskID, tagID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *TaskTagRepository) DeleteTaskTagsByTaskID(taskID uint) error {
	return r.db.Where("task_id = ?", taskID).Delete(&domain.TaskTag{}).Error
}
