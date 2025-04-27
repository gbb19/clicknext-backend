package postgres

import (
	"clicknext-backend/internal/domain"

	"gorm.io/gorm"
)

type BoardRepository struct {
	db *gorm.DB
}

// NewBoardRepository คือ constructor
func NewBoardRepository(db *gorm.DB) *BoardRepository {
	return &BoardRepository{db: db}
}

// CreateBoard สร้าง board ใหม่
func (r *BoardRepository) CreateBoard(board *domain.Board) error {
	return r.db.Create(board).Error
}

// GetBoardByID หา board จาก board ID
func (r *BoardRepository) GetBoardByID(id uint) (*domain.Board, error) {
	var board domain.Board
	if err := r.db.First(&board, id).Error; err != nil {
		return nil, err
	}
	return &board, nil
}

// ListBoards ดึง board list ทั้งหมด
func (r *BoardRepository) ListBoards() ([]*domain.Board, error) {
	var boards []*domain.Board
	if err := r.db.Find(&boards).Error; err != nil {
		return nil, err
	}
	return boards, nil
}

// AddBoardMember เพิ่มสมาชิกลงใน board
func (r *BoardRepository) AddBoardMember(boardMember *domain.BoardMember) error {
	return r.db.Create(boardMember).Error
}

// เพิ่มฟังก์ชันใหม่ใน BoardRepository
func (r *BoardRepository) GetBoardsByUserID(userID uint) ([]*domain.Board, error) {
	var boards []*domain.Board
	if err := r.db.Where("created_by = ?", userID).Find(&boards).Error; err != nil {
		return nil, err
	}
	return boards, nil
}

// เพิ่มฟังก์ชันใหม่ใน BoardRepository
func (r *BoardRepository) GetBoardsJoinedByUserID(userID uint) ([]*domain.Board, error) {
	var boards []*domain.Board
	if err := r.db.Joins("JOIN board_members ON boards.board_id = board_members.board_id").
		Where("board_members.user_id = ? AND boards.created_by != ?", userID, userID).
		Find(&boards).Error; err != nil {
		return nil, err
	}
	return boards, nil
}
