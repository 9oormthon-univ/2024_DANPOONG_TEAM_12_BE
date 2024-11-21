package contents

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"gorm.io/gorm"
)

type ContentsRepository struct {
	DB *gorm.DB
}

func SetContentsRepository(DB *gorm.DB) *ContentsRepository {
	u := &ContentsRepository{
		DB: DB,
	}
	return u
}

func (r *ContentsRepository) GetAllContents() ([]types.Content, error) {
	var contents []types.Content
	result := r.DB.Find(&contents)
	if result.Error != nil {
		return nil, result.Error
	}
	return contents, nil
}

func (r *ContentsRepository) GetContentByID(contentID int) (*types.Content, error) {
	var content types.Content
	result := *r.DB.First(&content, contentID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &content, nil
}
