package regions

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/gin-gonic/gin"
)

type RegionsController struct {
	regionsService types.RegionsService
}

func SetRegionsController(api *gin.RouterGroup, service types.RegionsService) *RegionsController {
	r := &RegionsController{
		regionsService: service,
	}
	// 핸들러 등록
	return r
}
