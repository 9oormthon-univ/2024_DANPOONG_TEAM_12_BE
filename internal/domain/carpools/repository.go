package carpools

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"gorm.io/gorm"
	"log"
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
func (repository *CarpoolsRepository) findByLocation(request types.GetCarpoolPostRequestDTO) ([]types.CarpoolPostResponseDTO, error) {
	var carpools []types.CarpoolPostResponseDTO

	// 조건에 맞는 카풀 게시글 조회
	query := repository.DB.Table("carpool").
		Where("start_location LIKE ?", "%"+request.StartLocation+"%").
		Where("end_location LIKE ?", "%"+request.EndLocation+"%").
		Where("status = ?", "active").
		Order("date ASC, start_time ASC"). // 날짜와 시간순으로 정렬
		Find(&carpools)

	if query.Error != nil {
		return nil, query.Error
	}
	return carpools, nil
}

// 내가 작성한 카풀 게시글 조회
func (repository *CarpoolsRepository) findByUser(request types.GetUserCarpoolPostRequestDTO) ([]types.CarpoolPostResponseDTO, error) {
	var carpools []types.CarpoolPostResponseDTO

	query := repository.DB.Table("carpool").
		Where("user_id = ?", request.UserID).
		Where("status = ?", "active").
		Order("created_at DESC").
		Scan(&carpools)

	// 에러 처리
	if query.Error != nil {
		log.Printf("Error fetching posts for user %d: %v", request.UserID, query.Error)
		return nil, query.Error
	}

	// 데이터가 없는 경우 처리
	if len(carpools) == 0 {
		log.Printf("No carpool posts found for user %d", request.UserID)
		return nil, nil // 데이터가 없을 경우 nil 반환
	}

	return carpools, nil
}
