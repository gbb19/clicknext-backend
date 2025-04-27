package postgres

import (
	"clicknext-backend/internal/domain"

	"gorm.io/gorm"
)

type ColumnRepository struct {
	db *gorm.DB
}

// NewColumnRepository returns a new ColumnRepository instance
func NewColumnRepository(db *gorm.DB) *ColumnRepository {
	return &ColumnRepository{db: db}
}

// CreateColumn creates a new column
func (r *ColumnRepository) CreateColumn(column *domain.Column) error {
	// Start a transaction
	tx := r.db.Begin()

	// Find max position for the given board
	var maxPosition struct {
		MaxPos int
	}
	if err := tx.Model(&domain.Column{}).
		Select("COALESCE(MAX(position), -1) as max_pos").
		Where("board_id = ?", column.BoardID).
		Scan(&maxPosition).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Set the position to max + 1 (or 0 if there are no columns yet)
	column.Position = maxPosition.MaxPos + 1

	// Create the column
	if err := tx.Create(column).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// GetColumnByID retrieves a column by its ID
func (r *ColumnRepository) GetColumnByID(id uint) (*domain.Column, error) {
	var column domain.Column
	if err := r.db.First(&column, id).Error; err != nil {
		return nil, err
	}
	return &column, nil
}

// GetColumnsByBoardID gets all columns for a board
func (r *ColumnRepository) GetColumnsByBoardID(boardID uint) ([]*domain.Column, error) {
	var columns []*domain.Column
	if err := r.db.Where("board_id = ?", boardID).Order("position ASC").Find(&columns).Error; err != nil {
		return nil, err
	}
	return columns, nil
}

// UpdateColumn updates a column
func (r *ColumnRepository) UpdateColumn(column *domain.Column) error {
	return r.db.Save(column).Error
}

// DeleteColumn deletes a column
func (r *ColumnRepository) DeleteColumn(id uint) error {
	return r.db.Delete(&domain.Column{}, id).Error
}

// UpdateColumnPosition updates the position of a column and adjusts other columns' positions
func (r *ColumnRepository) UpdateColumnPosition(columnID uint, newPosition int, boardID uint) error {
	// Start a transaction
	tx := r.db.Begin()

	// Find the column to update
	var column domain.Column
	if err := tx.First(&column, columnID).Error; err != nil {
		tx.Rollback()
		return err
	}

	oldPosition := column.Position

	// If position didn't change, do nothing
	if oldPosition == newPosition {
		tx.Rollback()
		return nil
	}

	// Update positions of other columns
	if oldPosition < newPosition {
		// Moving right: decrement positions of columns between old+1 and new position
		if err := tx.Model(&domain.Column{}).
			Where("board_id = ? AND position > ? AND position <= ?", boardID, oldPosition, newPosition).
			Update("position", gorm.Expr("position - 1")).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		// Moving left: increment positions of columns between new and old-1 position
		if err := tx.Model(&domain.Column{}).
			Where("board_id = ? AND position >= ? AND position < ?", boardID, newPosition, oldPosition).
			Update("position", gorm.Expr("position + 1")).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Update the position of the target column
	column.Position = newPosition
	if err := tx.Save(&column).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
