package courses

import (
	"net/http"

	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/util"
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
	api.GET("/courses/ai", c.RecommendCourses)
	api.POST("/courses", c.PostMyCourses)
	return c
}

// RecommendCourses godoc
// @Summary 코스 추천
// @Description 코스 추천이 성공적으로 완료되었습니다.
// @Tags Courses
// @Accept json
// @Produce json
// @Param request body types.RecommendCourseReq true "추천 요청 데이터"
// @Success 200 {object} util.ResponseDTO{data=types.Course} "성공 응답 데이터"
// @Failure 400 {object} util.ResponseDTO "에러 응답 데이터"
// @Router /courses/ai [post]
func (c *CoursesController) RecommendCourses(ctx *gin.Context) {
	var req types.RecommendCourseReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err.Error()))
		return
	}

	course, err := c.coursesService.RecommendCourses(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, util.SuccessResponse("코스 추천 성공", course))
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
