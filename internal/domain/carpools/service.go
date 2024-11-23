package carpools

import (
	"errors"
	"fmt"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"time"
)

type CarpoolsService struct {
	carpoolsRepository *CarpoolsRepository
}

func SetCarpoolsService(repository *CarpoolsRepository) types.CarpoolsService {
	r := &CarpoolsService{
		carpoolsRepository: repository,
	}
	return r
}

// 카풀 게시글 좋아요순 조회
func (service *CarpoolsService) GetTopLikedCarpools(limit int) ([]types.CarpoolTopLikesResponseDTO, error) {
	return service.carpoolsRepository.GetTopLikedCarpools(limit)
}

// 카풀 게시글 생성
func (service *CarpoolsService) CreateCarpoolsPost(request types.CreateCarpoolPostRequestDTO) error {
	// 입력 데이터 유효성 검사
	if _, err := time.Parse("2006-01-02", request.Date); err != nil {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
		return errors.New("Invalid date format. Please use YYYY-MM-DD.")
	}

	if _, err := time.Parse("15:04:05", request.StartTime); err != nil {
		fmt.Println("Invalid start time format. Please use HH:MM:SS.")
		return errors.New("Invalid start time format. Please use HH:MM:SS.")
	}

	carpools := types.Carpool{
		Title:         request.Title,
		ImageURL:      request.ImageURL,
		Details:       request.Details,
		UserID:        request.UserID,
		StartLocation: request.StartLocation,
		EndLocation:   request.EndLocation,
		Date:          request.Date,
		StartTime:     request.StartTime,
	}

	if err := service.carpoolsRepository.SaveCarpoolPost(carpools); err != nil {
		return errors.New("failed to save matching post")
	}
	return nil
}
