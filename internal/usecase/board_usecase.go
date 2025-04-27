package usecase

import (
	"clicknext-backend/internal/domain"
)

type BoardRepository interface {
	CreateBoard(board *domain.Board) error
	GetBoardByID(id uint) (*domain.Board, error)
	ListBoards() ([]*domain.Board, error)
	GetBoardsByUserID(userID uint) ([]*domain.Board, error)
	GetBoardsJoinedByUserID(userID uint) ([]*domain.Board, error)
}

type BoardUsecase struct {
	repo BoardRepository
}

func NewBoardUsecase(repo BoardRepository) *BoardUsecase {
	return &BoardUsecase{repo: repo}
}

func (uc *BoardUsecase) CreateBoard(board *domain.Board) error {
	return uc.repo.CreateBoard(board)
}

func (uc *BoardUsecase) GetBoardByID(id uint) (*domain.Board, error) {
	return uc.repo.GetBoardByID(id)
}

func (uc *BoardUsecase) ListBoards() ([]*domain.Board, error) {
	return uc.repo.ListBoards()
}

func (uc *BoardUsecase) GetBoardsByUserID(userID uint) ([]*domain.Board, error) {
	return uc.repo.GetBoardsByUserID(userID)
}

// เพิ่มฟังก์ชันใน BoardUsecase
func (uc *BoardUsecase) GetBoardsJoinedByUserID(userID uint) ([]*domain.Board, error) {
	return uc.repo.GetBoardsJoinedByUserID(userID)
}
