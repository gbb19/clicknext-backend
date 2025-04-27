package usecase

import (
	"clicknext-backend/internal/domain"
)

type BoardMemberRepository interface {
	GetBoardMemberByBoardID(boardID uint) ([]*domain.BoardMember, error)
	AddBoardMember(boardID uint, userID uint) error
}

type BoardMemberUsecase struct {
	repo BoardMemberRepository
}

func NewBoardMemberUsecase(repo BoardMemberRepository) *BoardMemberUsecase {
	return &BoardMemberUsecase{repo: repo}
}

func (uc *BoardMemberUsecase) GetBoardMemberByBoardID(boardID uint) ([]*domain.BoardMember, error) {
	return uc.repo.GetBoardMemberByBoardID(boardID)
}
func (uc *BoardMemberUsecase) AddBoardMember(boardID uint, userID uint) error {
	return uc.repo.AddBoardMember(boardID, userID)
}
