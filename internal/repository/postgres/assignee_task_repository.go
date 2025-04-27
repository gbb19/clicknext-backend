package postgres

import (
	"clicknext-backend/internal/domain"

	"gorm.io/gorm"
)

type AssigneeTaskRepository struct {
	db *gorm.DB
}

func NewAssigneeTaskRepository(db *gorm.DB) *AssigneeTaskRepository {
	return &AssigneeTaskRepository{db: db}
}

func (r *AssigneeTaskRepository) AssignTask(assigneeTask *domain.AssigneeTask) error {
	// Check if assignment already exists
	var existingAssignment domain.AssigneeTask
	if err := r.db.Where("assignee_id = ? AND task_id = ?",
		assigneeTask.AssigneeID, assigneeTask.TaskID).
		First(&existingAssignment).Error; err == nil {
		// Assignment already exists
		return nil
	}

	return r.db.Create(assigneeTask).Error
}

func (r *AssigneeTaskRepository) UnassignTask(assigneeID, taskID uint) error {
	return r.db.Where("assignee_id = ? AND task_id = ?", assigneeID, taskID).
		Delete(&domain.AssigneeTask{}).Error
}

func (r *AssigneeTaskRepository) GetAssigneesByTaskID(taskID uint) ([]*domain.AssigneeTask, error) {
	var assignees []*domain.AssigneeTask
	if err := r.db.
		Joins("JOIN users ON assignee_tasks.assignee_id = users.user_id").
		Joins("JOIN tasks ON assignee_tasks.task_id = tasks.task_id").
		Where("assignee_tasks.task_id = ?", taskID).
		Find(&assignees).Error; err != nil {
		return nil, err
	}
	return assignees, nil
}

func (r *AssigneeTaskRepository) GetTasksByAssigneeID(assigneeID uint) ([]*domain.AssigneeTask, error) {
	var tasks []*domain.AssigneeTask
	if err := r.db.
		Joins("JOIN tasks ON assignee_tasks.task_id = tasks.task_id").
		Joins("JOIN users ON assignee_tasks.assignee_id = users.user_id").
		Where("assignee_tasks.assignee_id = ?", assigneeID).
		Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
