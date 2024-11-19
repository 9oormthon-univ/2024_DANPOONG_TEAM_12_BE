package courses

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/gin-gonic/gin"
)

type CoursesController struct {
	coursesService types.CoursesService
}

func SetCoursesController(api *gin.RouterGroup, service types.CoursesService) *CoursesController {
	c := &CoursesController{
		coursesService: service,
	}
	// 핸들러 등록
	return c
}
