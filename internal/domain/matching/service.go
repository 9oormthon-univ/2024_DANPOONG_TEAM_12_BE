package matching

import (
	"errors"
	"fmt"
	"time"

	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/users"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
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

// 모든 게시글을 가져와서 id, details, categories 추출 --- 사용자 위치는 아직 추출 x
func (service *matchingService) GetPostsForAI(page int, pageSize int) ([]*types.MatchingDetailForAI, error) {

	var results []*types.MatchingDetailForAI = make([]*types.MatchingDetailForAI, 0)

	posts, err := service.matchingRepository.GetAllMatchingPosts(page, pageSize)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		result := &types.MatchingDetailForAI{
			MatchingID: string(post.MatchingID),
			Details:    post.Details,
			Categories: []string{},
		}
		result.Categories = append(result.Categories, result.Categories...)
		results = append(results, result)
	}

	return results, nil
}

func (service *matchingService) GetExampleMatchingPosts() ([]*types.MatchingDetailForAI, error) {
	// 예시 데이터 생성
	examplePosts := []*types.MatchingDetailForAI{
		{
			MatchingID: string(1),
			Details:    "친구와 함께하는 주말 등산 모임입니다. 자연을 사랑하고 활발한 활동을 즐기는 분들을 모집합니다.",
			Categories: []string{"친구", "등산", "활동적"},
		},
		{
			MatchingID: string(2),
			Details:    "내향인 분들을 위한 조용한 독서 클럽입니다. 편안한 분위기에서 다양한 책을 함께 읽고 토론해요.",
			Categories: []string{"내향인", "독서", "클럽"},
		},
		{
			MatchingID: string(3),
			Details:    "프로그래밍과 테크놀로지에 관심 있는 분들을 위한 워크샵을 개최합니다. 최신 기술 동향을 함께 논의해요.",
			Categories: []string{"프로그래밍", "테크놀로지", "워크샵"},
		},
	}

	// 예시 데이터를 반환
	return examplePosts, nil
}
