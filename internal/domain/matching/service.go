package matching

import (
	"errors"
	"fmt"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/users"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"time"
)

type matchingService struct {
	matchingRepository *MatchingRepository
}

func SetMatchingService(repository *MatchingRepository) *matchingService {
	return &matchingService{
		matchingRepository: repository,
	}
}

// 매칭 게시글 좋아요순 조회
func (service *matchingService) GetTopLikeMatchingPosts(limit int) ([]types.MatchingTopLikesResponseDTO, error) {
	return service.matchingRepository.GetTopLikedMatchingPosts(limit)
}

// 매칭 게시글 생성
func (service *matchingService) CreateMatchingPost(request types.CreateMatchingPostRequestDTO) error {

	// 입력 데이터 유효성 검사
	if _, err := time.Parse("2006-01-02", request.Date); err != nil {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
		return errors.New("Invalid date format. Please use YYYY-MM-DD.")
	}

	if _, err := time.Parse("15:04:05", request.StartTime); err != nil {
		fmt.Println("Invalid start time format. Please use HH:MM:SS.")
		return errors.New("Invalid start time format. Please use HH:MM:SS.")
	}

	if _, err := time.Parse("15:04:05", request.EndTime); err != nil {
		fmt.Println("Invalid end time format. Please use HH:MM:SS.")
		return errors.New("Invalid end time format. Please use HH:MM:SS.")
	}

	matching := types.Matching{
		Title:        request.Title,
		ImageURL:     request.ImageURL,
		Details:      request.Details,
		UserNickname: request.UserNickname,
		UserID:       request.UserID,
		Destination:  request.Destination,
		Date:         request.Date,
		StartTime:    request.StartTime,
		EndTime:      request.EndTime,
	}

	if err := service.matchingRepository.SaveMatchingPost(matching); err != nil {
		return errors.New("failed to save matching post")
	}
	return nil
}

// 사용자 매칭 게시글 목록 조회
func (service *matchingService) GetUserMatchingPosts(userID int64) ([]types.MatchingPostResponseDTO, error) {
	// 사용자 존재 여부 확인
	userRepo := users.UsersRepository{DB: service.matchingRepository.DB} // DB 인스턴스를 전달하여 UsersRepository 생성
	user, err := userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user with ID %d: %v", userID, err)
	}
	if user == nil {
		return nil, fmt.Errorf("user with ID %d does not exist", userID)
	}
	return service.matchingRepository.GetUserMatchingPosts(userID)
}

// 매칭 지원하기
func (service *matchingService) CreateMatchingApplication(request types.MatchingApplicationRequestDTO, matchingID int64) (*types.MatchingApplicationResponseDTO, error) {
	// 매칭 게시글 존재 여부 확인
	matching, err := service.matchingRepository.GetByID(matchingID)
	if err != nil {
		return nil, fmt.Errorf("matching not found: %v", err)
	}
	if matching == nil {
		return nil, fmt.Errorf("matching with ID %d does not exist", matchingID)
	}

	// 사용자 존재 여부 확인
	userRepo := users.UsersRepository{DB: service.matchingRepository.DB} // DB 인스턴스를 전달하여 UsersRepository 생성
	user, err := userRepo.GetByID(request.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user with ID %d: %v", request.UserID, err)
	}
	if user == nil {
		return nil, fmt.Errorf("user with ID %d does not exist", request.UserID)
	}

	//매칭 지원서 생성
	application := &types.MatchingApplication{
		UserID:      request.UserID,
		MatchingID:  matchingID,
		Description: request.Description,
	}

	// 데이터 베이스 저장
	if err := service.matchingRepository.CreateMatchingApplication(application); err != nil {
		return nil, err
	}

	// 응답 DTO 생성
	response := &types.MatchingApplicationResponseDTO{
		ApplicationID: application.ApplicationID,
		MatchingID:    application.MatchingID,
		UserID:        application.UserID,
		Description:   application.Description,
		AppliedAt:     application.AppliedAt,
		User:          *user,
	}

	return response, nil
}
