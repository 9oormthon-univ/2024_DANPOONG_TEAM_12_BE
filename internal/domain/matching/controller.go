package matching

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/gin-gonic/gin"
)

type MatchingController struct {
	matchingService types.MatchingService
}

func SetMatchingController(api *gin.RouterGroup, service types.MatchingService) *MatchingController {
	m := &MatchingController{
		matchingService: service,
	}
	// 핸들러 등록
	return m
}
