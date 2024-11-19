package ai

import "github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"

type AIService struct {
	aiRepository *AIRepository
}

func SetAIService(repository *AIRepository) types.AIService {
	r := &AIService{
		aiRepository: repository,
	}
	return r
}
