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

// 카풀 게시글 조회(출발 위치, 목적지 위치 기반)
func (repository *CarpoolsRepository) findByLocation(request types.GetCarpoolPostRequestDTO) ([]types.Carpool, error) {
	var carpoolList []types.Carpool

	// 조건에 맞는 카풀 게시글 조회
	query := repository.DB.Table("carpool").
		Where("start_location LIKE ?", "%"+request.StartLocation+"%").
		Where("end_location LIKE ?", "%"+request.EndLocation+"%").
		Where("status = ?", "active").
		Order("date ASC, start_time ASC"). // 날짜와 시간순으로 정렬
		Find(&carpoolList)

	if query.Error != nil {
		return nil, query.Error
	}
	return carpoolList, nil
}
