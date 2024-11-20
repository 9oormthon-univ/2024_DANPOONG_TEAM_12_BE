package carpools

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/gin-gonic/gin"
)

type CarpoolsController struct {
	carpoolsService types.CarpoolsService
}

func SetCarpoolsController(api *gin.RouterGroup, service types.CarpoolsService) *CarpoolsController {
	c := &CarpoolsController{
		carpoolsService: service,
	}
	// 핸들러 등록
	return c
}

func (c *CarpoolsController) PostCourseAI(ctx *gin.Context) {

}
