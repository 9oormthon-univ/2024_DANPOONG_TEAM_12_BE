package regions

import "github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"

type regionsService struct {
	regionsRepository *RegionsRepository
}

func SetRegionsService(repository *RegionsRepository) types.RegionsService {
	r := &regionsService{
		regionsRepository: repository,
	}
	return r
}
