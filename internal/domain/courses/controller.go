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
	api.POST("/courses", c.PostMyCourses)
	return c
}

func (c *CoursesController) RecommendCourses(ctx *gin.Context) {
	var req types.RecommendCourseReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"type":    "error",
			"message": err.Error(),
		})
		return
	}

	course, err := c.coursesService.RecommendCourses(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"type":    "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"type":    "success",
		"message": "코스 추천 성공",
		"data":    course,
	})

}

func (c *CoursesController) GetMyCourses(ctx *gin.Context) {

}

func (c *CoursesController) PostMyCourses(ctx *gin.Context) {

}

func (c *CoursesController) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "/courses api health check",
	})
}
