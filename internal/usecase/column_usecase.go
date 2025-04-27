package usecase

import (
	"clicknext-backend/internal/domain"
)

type ColumnRepository interface {
	CreateColumn(column *domain.Column) error
	GetColumnByID(id uint) (*domain.Column, error)
	GetColumnsByBoardID(boardID uint) ([]*domain.Column, error)
	UpdateColumn(column *domain.Column) error
	DeleteColumn(id uint) error
	UpdateColumnPosition(columnID uint, newPosition int, boardID uint) error
}

type ColumnUsecase struct {
	repo ColumnRepository
}

func NewColumnUsecase(repo ColumnRepository) *ColumnUsecase {
	return &ColumnUsecase{repo: repo}
}

func (uc *ColumnUsecase) CreateColumn(column *domain.Column) error {
	return uc.repo.CreateColumn(column)
}

func (uc *ColumnUsecase) GetColumnByID(id uint) (*domain.Column, error) {
	return uc.repo.GetColumnByID(id)
}

func (uc *ColumnUsecase) GetColumnsByBoardID(boardID uint) ([]*domain.Column, error) {
	return uc.repo.GetColumnsByBoardID(boardID)
}

func (uc *ColumnUsecase) UpdateColumn(column *domain.Column) error {
	return uc.repo.UpdateColumn(column)
}

func (uc *ColumnUsecase) DeleteColumn(id uint) error {
	return uc.repo.DeleteColumn(id)
}

func (uc *ColumnUsecase) UpdateColumnPosition(columnID uint, newPosition int, boardID uint) error {
	return uc.repo.UpdateColumnPosition(columnID, newPosition, boardID)
}
