package contents

import "github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"

type contentsService struct {
	contentsRepository ContentsRepository
}

func SetContentsService(repository ContentsRepository) types.ContentsService {
	return &contentsService{
		contentsRepository: repository,
	}
}

func (s *contentsService) GetAllContents() ([]types.Content, error) {
	return s.contentsRepository.GetAllContents()
}

func (s *contentsService) GetContentById(contentID int) (*types.Content, error) {
	return s.contentsRepository.GetContentByID(contentID)
}
