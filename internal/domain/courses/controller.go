package courses

import (
	"net/http"

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
	api.GET("/courses/health", c.Health)

	api.GET("/courses/me", c.GetMyCourses)
	api.POST("/courses/ai", c.RecommendCourses)
	return c
}

func (c *CoursesController) RecommendCourses(ctx *gin.Context) {
	// 관심사, 위치 -> ai 코스 추천
}

func (c *CoursesController) GetMyCourses(ctx *gin.Context) {

}

func (c *CoursesController) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "/courses api health check",
	})
}
