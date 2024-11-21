package carpools

import "github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"

type CarpoolsService struct {
	carpoolsRepository *CarpoolsRepository
}

func SetCarpoolsService(repository *CarpoolsRepository) types.CarpoolsService {
	r := &CarpoolsService{
		carpoolsRepository: repository,
	}
	return r
}

func (service *CarpoolsService) GetTopLikedCarpools(limit int) ([]types.CarpoolTopLikesResponseDTO, error) {
	return service.carpoolsRepository.GetTopLikedCarpools(limit)
}
