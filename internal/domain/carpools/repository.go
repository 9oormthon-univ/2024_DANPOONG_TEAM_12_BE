package carpools

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"gorm.io/gorm"
)

type CarpoolsRepository struct {
	DB *gorm.DB
}

func SetCarpoolsRepository(DB *gorm.DB) *CarpoolsRepository {
	r := &CarpoolsRepository{
		DB: DB,
	}
	return r
}

// 카풀 게시글 좋아요 순 조회
func (repository *CarpoolsRepository) GetTopLikedCarpools(limit int) ([]types.CarpoolTopLikesResponseDTO, error) {
	var carpools []types.CarpoolTopLikesResponseDTO
	result := repository.DB.Table("carpool").
		Select("carpool_id, title, details, start_location, end_location, start_time, likes").
		Where("status =?", "active").
		Order("likes DESC").
		Limit(limit).
		Scan(&carpools)

	if result.Error != nil {
		return nil, result.Error
	}
	return carpools, nil
}

// 카풀 게시글 생성
func (repository *CarpoolsRepository) SaveCarpoolPost(post types.Carpool) error {
	if err := repository.DB.Create(&post).Error; err != nil {
		return err
	}
	return nil
}
