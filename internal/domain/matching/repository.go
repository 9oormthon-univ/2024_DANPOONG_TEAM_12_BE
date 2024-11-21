package matching

import (
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
