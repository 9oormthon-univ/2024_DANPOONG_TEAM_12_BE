package ai

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/gin-gonic/gin"
)

type AIController struct {
	aiService types.AIService
}

func SetAIController(api *gin.RouterGroup, service types.AIService) *AIController {
	c := &AIController{
		aiService: service,
	}
	// 핸들러 등록
	return c
}
