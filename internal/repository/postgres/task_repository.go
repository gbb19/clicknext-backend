package postgres

import (
	"clicknext-backend/internal/domain"
	"log"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task *domain.Task) error {
	tx := r.db.Begin()

	// Find max position for the given column
	var maxPosition struct {
		MaxPos int
	}
	if err := tx.Model(&domain.Task{}).
		Select("COALESCE(MAX(position), -1) as max_pos").
		Where("column_id = ?", task.ColumnID).
		Scan(&maxPosition).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Set position to max + 1 (or 0 if no tasks yet)
	task.Position = maxPosition.MaxPos + 1

	// Create the task
	if err := tx.Create(task).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *TaskRepository) GetTaskByID(id uint) (*domain.Task, error) {
	var task domain.Task
	if err := r.db.
		Joins("JOIN users ON tasks.created_by = users.user_id").
		Joins("JOIN columns ON tasks.column_id = columns.column_id").
		Where("tasks.task_id = ?", id).
		First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) GetTasksByColumnID(columnID uint) ([]*domain.Task, error) {
	var tasks []*domain.Task
	if err := r.db.
		Joins("JOIN users ON tasks.created_by = users.user_id").
		Joins("JOIN columns ON tasks.column_id = columns.column_id").
		Where("tasks.column_id = ?", columnID).
		Order("tasks.position ASC").
		Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) UpdateTask(task *domain.Task) error {
	err := r.db.Save(task).Error
	if err != nil {
		log.Println("Error updating task:", err)
	}
	return err
}

func (r *TaskRepository) DeleteTask(id uint) error {
	return r.db.Delete(&domain.Task{}, id).Error
}

func (r *TaskRepository) UpdateTaskPosition(taskID uint, newPosition int, columnID uint) error {
	tx := r.db.Begin()

	// Find the task to update
	var task domain.Task
	if err := tx.First(&task, taskID).Error; err != nil {
		tx.Rollback()
		return err
	}

	oldPosition := task.Position

	// If position didn't change, do nothing
	if oldPosition == newPosition {
		tx.Rollback()
		return nil
	}

	// Update positions of other tasks
	if oldPosition < newPosition {
		// Moving right: decrement positions of tasks between old+1 and new position
		if err := tx.Model(&domain.Task{}).
			Where("column_id = ? AND position > ? AND position <= ?", columnID, oldPosition, newPosition).
			Update("position", gorm.Expr("position - 1")).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		// Moving left: increment positions of tasks between new and old-1 position
		if err := tx.Model(&domain.Task{}).
			Where("column_id = ? AND position >= ? AND position < ?", columnID, newPosition, oldPosition).
			Update("position", gorm.Expr("position + 1")).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Update the position of the target task
	task.Position = newPosition
	if err := tx.Save(&task).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
