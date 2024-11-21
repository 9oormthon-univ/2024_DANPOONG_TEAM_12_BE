package ai

import (
	"net/http"

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
	api.POST("/ai/test/data", c.GenerateTrainingData)
	return c
}

func (a *AIController) GenerateTrainingData(ctx *gin.Context) {
	err := a.aiService.GenerateTrainingData()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"type":    "error",
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"type": "success",
	})
}
