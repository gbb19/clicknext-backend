package postgres

import (
	"clicknext-backend/internal/domain"

	"gorm.io/gorm"
)

type InviteRepository struct {
	db *gorm.DB
}

// NewInviteRepository คือ constructor
func NewInviteRepository(db *gorm.DB) *InviteRepository {
	return &InviteRepository{db: db}
}

// CreateInvite สร้าง invite ใหม่
func (r *InviteRepository) CreateInvite(invite *domain.Invite) error {
	return r.db.Create(invite).Error
}

// GetInviteByID หา invite จาก id
func (r *InviteRepository) GetInviteByID(id uint) (*domain.Invite, error) {
	var invite domain.Invite
	if err := r.db.First(&invite, id).Error; err != nil {
		return nil, err
	}
	return &invite, nil
}

// GetInvitesByBoardID หา invites ตาม Board ID
func (r *InviteRepository) GetInvitesByBoardID(boardID uint) ([]*domain.Invite, error) {
	var invites []*domain.Invite
	if err := r.db.Where("board_id = ?", boardID).Find(&invites).Error; err != nil {
		return nil, err
	}
	return invites, nil
}

// UpdateInvite อัปเดต status ของ invite
func (r *InviteRepository) UpdateInvite(invite *domain.Invite) error {
	return r.db.Save(invite).Error
}

func (r *InviteRepository) UpdateInviteStatus(inviteID uint, status string) error {
	return r.db.Model(&domain.Invite{}).Where("invite_id = ?", inviteID).Update("status", status).Error
}

// DeleteInvite ลบ invite
func (r *InviteRepository) DeleteInvite(id uint) error {
	return r.db.Delete(&domain.Invite{}, id).Error
}
