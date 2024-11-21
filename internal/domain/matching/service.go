package matching

import "github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"

type matchingService struct {
	matchingRepository *MatchingRepository
}

func SetMatchingService(repository *MatchingRepository) *matchingService {
	return &matchingService{
		matchingRepository: repository,
	}
}

func (s *matchingService) GetTopLikeMatchingPosts(limit int) ([]types.MatchingTopLikesResponseDTO, error) {
	return s.matchingRepository.GetTopLikedMatchingPosts(limit)
}
