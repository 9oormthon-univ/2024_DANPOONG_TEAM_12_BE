package matching

import (
	"errors"
	"fmt"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"gorm.io/gorm"
	"log"
)

type MatchingRepository struct {
	DB *gorm.DB
}

func SetMatchingRepository(DB *gorm.DB) *MatchingRepository {
	r := &MatchingRepository{
		DB: DB,
	}
	return r
}

func (r *MatchingRepository) GetTopLikedMatchingPosts(limit int) ([]types.MatchingTopLikesResponseDTO, error) {
	var posts []types.MatchingTopLikesResponseDTO
	result := r.DB.Table("matching").
		Select("matching_id, title, image_url, details, likes AS likes_count").
		Where("status = ?", "active").
		Order("likes DESC").
		Limit(limit).
		Scan(&posts)

	// 디버깅 로그 추가
	if result.Error != nil {
		log.Printf("Error fetching posts: %v", result.Error)
		return nil, result.Error
	}

	// 데이터가 없을 때
	if len(posts) == 0 {
		log.Println("No matching posts found")
		return nil, result.Error
	}
	return posts, nil
}

// 매칭 게시글 저장
func (repository *MatchingRepository) SaveMatchingPost(post types.Matching) error {
	if err := repository.DB.Create(&post).Error; err != nil {
		return err
	}
	return nil
}

// 내가 작성한 매칭 게시글 조회
func (repository *MatchingRepository) GetUserMatchingPosts(userID int64) ([]types.MatchingPostResponseDTO, error) {
	var posts []types.MatchingPostResponseDTO

	result := repository.DB.Table("matching").
		Select("matching_id, title, image_url, details, user_nickname, destination, "+
			"DATE_FORMAT(date, '%Y-%m-%d') AS date, "+
			"DATE_FORMAT(start_time, '%H:%i:%s') AS start_time, "+
			"DATE_FORMAT(end_time, '%H:%i:%s') AS end_time, likes").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Scan(&posts)

	// 에러 처리
	if result.Error != nil {
		log.Printf("Error fetching posts for user %d: %v", userID, result.Error)
		return nil, result.Error
	}

	// 데이터가 없는 경우 처리
	if len(posts) == 0 {
		log.Printf("No matching posts found for user %d", userID)
		return nil, nil // 데이터가 없을 경우 nil 반환
	}

	return posts, nil
}

// 매칭 지원서 작성
func (repository *MatchingRepository) CreateMatchingApplication(application *types.MatchingApplication) error {
	if err := repository.DB.Create(application).Error; err != nil {
		return err
	}
	return nil
}

// GetByID - 매칭 ID로 매칭 게시글을 조회
func (r *MatchingRepository) GetByID(matchingID int64) (*types.Matching, error) {
	var matching types.Matching

	// 매칭 게시글 조회
	if err := r.DB.Where("matching_id = ?", matchingID).First(&matching).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("matching with ID %d not found", matchingID)
		}
		return nil, err
	}

	return &matching, nil
}
