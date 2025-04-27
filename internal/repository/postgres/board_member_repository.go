package postgres

import (
	"clicknext-backend/internal/domain"

	"gorm.io/gorm"
)

type BoardMemberRepository struct {
	db *gorm.DB
}

// NewBoardMemberRepository คือ constructor
func NewBoardMemberRepository(db *gorm.DB) *BoardMemberRepository {
	return &BoardMemberRepository{db: db}
}

// GetBoardMemberByBoardID หาสมาชิกบอร์ดจาก BoardID
func (r *BoardMemberRepository) GetBoardMemberByBoardID(boardID uint) ([]*domain.BoardMember, error) {
	var members []*domain.BoardMember
	if err := r.db.Where("board_id = ?", boardID).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *BoardMemberRepository) AddBoardMember(boardID uint, userID uint) error {
	boardMember := &domain.BoardMember{
		BoardID: boardID,
		UserID:  userID,
	}
	return r.db.Create(boardMember).Error
}

// เพิ่มฟังก์ชันใน BoardMemberRepository
func (r *BoardMemberRepository) GetBoardsByUserID(userID uint) ([]*domain.Board, error) {
	var boards []*domain.Board
	if err := r.db.Table("boards").
		Joins("JOIN board_members ON boards.board_id = board_members.board_id").
		Where("board_members.user_id = ?", userID).
		Find(&boards).Error; err != nil {
		return nil, err
	}
	return boards, nil
}
