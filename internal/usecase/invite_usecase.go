package usecase

import (
	"clicknext-backend/internal/domain"
)

type InviteRepository interface {
	CreateInvite(invite *domain.Invite) error
	GetInviteByID(id uint) (*domain.Invite, error)
	GetInvitesByBoardID(boardID uint) ([]*domain.Invite, error)
	UpdateInvite(invite *domain.Invite) error
	DeleteInvite(id uint) error
	UpdateInviteStatus(inviteID uint, status string) error
}

type InviteUsecase struct {
	repo            InviteRepository
	boardMemberRepo BoardMemberRepository
}

func NewInviteUsecase(repo InviteRepository, boardMemberRepo BoardMemberRepository) *InviteUsecase {
	return &InviteUsecase{repo: repo, boardMemberRepo: boardMemberRepo}
}

func (uc *InviteUsecase) CreateInvite(invite *domain.Invite) error {
	return uc.repo.CreateInvite(invite)
}

func (uc *InviteUsecase) GetInviteByID(id uint) (*domain.Invite, error) {
	return uc.repo.GetInviteByID(id)
}

func (uc *InviteUsecase) GetInvitesByBoardID(boardID uint) ([]*domain.Invite, error) {
	return uc.repo.GetInvitesByBoardID(boardID)
}

func (uc *InviteUsecase) UpdateInvite(invite *domain.Invite) error {
	return uc.repo.UpdateInvite(invite)
}

func (uc *InviteUsecase) DeleteInvite(id uint) error {
	return uc.repo.DeleteInvite(id)
}

func (uc *InviteUsecase) AcceptInvite(inviteID uint, userID uint) error {
	// อัปเดตสถานะของคำเชิญเป็น "accepted"
	if err := uc.repo.UpdateInviteStatus(inviteID, "accepted"); err != nil {
		return err
	}

	// ดึงข้อมูลของคำเชิญเพื่อหาข้อมูล Board
	invite, err := uc.repo.GetInviteByID(inviteID)
	if err != nil {
		return err
	}

	// เพิ่มสมาชิกใหม่ใน Board
	if err := uc.boardMemberRepo.AddBoardMember(invite.BoardID, userID); err != nil {
		return err
	}

	return nil
}
