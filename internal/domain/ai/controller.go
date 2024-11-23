package ai

import (
	"net/http"

	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/util"
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
	api.POST("/ai/test/course", c.RecommendCoursesTest)
	api.POST("/ai/test/matching", c.RecommendMatchingPostTest)
	return c
}

func (a *AIController) RecommendCoursesTest(ctx *gin.Context) {
	var req types.RecommendCourseReq
	// 관심사와 위치를 받아온다
	// 관심사는 Tour API와 맵핑은 안 되므로
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
		ctx.JSON(http.StatusInternalServerError, gin.H{ // 상태 코드 수정
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

func (a *AIController) RecommendMatchingPostTest(ctx *gin.Context) {
	location := "서울"
	interests := []string{"코딩", "기술"}

	posts, err := a.aiService.RecommendMatchingPost(1, 5, location, interests)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, posts)
}
