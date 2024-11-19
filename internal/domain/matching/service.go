package matching

import "github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"

type matchingService struct {
	matchingRepository *MatchingRepository
}

func SetMatchingService(repository *MatchingRepository) types.MatchingService {
	r := &matchingService{
		matchingRepository: repository,
	}
	return r
}
