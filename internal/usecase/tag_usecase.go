package usecase

import (
	"clicknext-backend/internal/domain"
)

type TagRepository interface {
	CreateTag(tag *domain.Tag) error
	GetTagByID(id uint) (*domain.Tag, error)
	GetAllTags() ([]*domain.Tag, error)
	UpdateTag(tag *domain.Tag) error
	DeleteTag(id uint) error
	GetTagsByCreatedBy(userID uint) ([]*domain.Tag, error)
	FindTagByName(name string) (*domain.Tag, error)
}

type TagUsecase struct {
	repo TagRepository
}

func NewTagUsecase(repo TagRepository) *TagUsecase {
	return &TagUsecase{repo: repo}
}

func (uc *TagUsecase) CreateTag(tag *domain.Tag) error {
	return uc.repo.CreateTag(tag)
}

func (uc *TagUsecase) GetTagByID(id uint) (*domain.Tag, error) {
	return uc.repo.GetTagByID(id)
}

func (uc *TagUsecase) GetAllTags() ([]*domain.Tag, error) {
	return uc.repo.GetAllTags()
}

func (uc *TagUsecase) UpdateTag(tag *domain.Tag) error {
	return uc.repo.UpdateTag(tag)
}

func (uc *TagUsecase) DeleteTag(id uint) error {
	return uc.repo.DeleteTag(id)
}

func (uc *TagUsecase) GetTagsByCreatedBy(userID uint) ([]*domain.Tag, error) {
	return uc.repo.GetTagsByCreatedBy(userID)
}

func (uc *TagUsecase) FindTagByName(name string) (*domain.Tag, error) {
	return uc.repo.FindTagByName(name)
}
