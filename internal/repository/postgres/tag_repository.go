package postgres

import (
	"clicknext-backend/internal/domain"

	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) CreateTag(tag *domain.Tag) error {
	return r.db.Create(tag).Error
}

func (r *TagRepository) GetTagByID(id uint) (*domain.Tag, error) {
	var tag domain.Tag
	if err := r.db.Preload("CreatedByUser").First(&tag, id).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *TagRepository) GetAllTags() ([]*domain.Tag, error) {
	var tags []*domain.Tag
	if err := r.db.Preload("CreatedByUser").Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *TagRepository) UpdateTag(tag *domain.Tag) error {
	return r.db.Save(tag).Error
}

func (r *TagRepository) DeleteTag(id uint) error {
	return r.db.Delete(&domain.Tag{}, id).Error
}

func (r *TagRepository) GetTagsByCreatedBy(userID uint) ([]*domain.Tag, error) {
	var tags []*domain.Tag
	if err := r.db.Where("created_by = ?", userID).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *TagRepository) FindTagByName(name string) (*domain.Tag, error) {
	var tag domain.Tag
	if err := r.db.Where("name = ?", name).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}
