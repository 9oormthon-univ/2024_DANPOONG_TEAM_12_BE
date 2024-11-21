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
	api.POST("/ai/test/train", c.TrainingModel)
	api.POST("/ai/test/course", c.RecommendCoursesTest)
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

func (a *AIController) TrainingModel(ctx *gin.Context) {
	err := a.aiService.RequestFineTuning()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"type":    "error",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"type": "success",
	})
}

func (a *AIController) RecommendCoursesTest(ctx *gin.Context) {
	var req types.RecommendCourseRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"type":    "error",
		})
		return
	}

	result, err := a.aiService.RecommendCourses(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"type":    "error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "코스 추천 성공",
		"type":    "success",
		"data":    result,
	})

}
