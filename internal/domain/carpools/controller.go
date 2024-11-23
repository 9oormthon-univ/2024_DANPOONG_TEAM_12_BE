package carpools

import (
	"net/http"
	"strconv"

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
	// 카풀 게시글 좋아요 순 조회
	api.GET("/carpools/posts/sorted-by-likes", c.GetTopLikedCarpools)
	// 카풀 게시글 생성
	api.POST("/carpools/posts", c.CreateCarpoolPost)

	return c
}

// 카풀 게시글 좋아요순 조회
func (controller *CarpoolsController) GetTopLikedCarpools(ctx *gin.Context) {
	limitStr := ctx.DefaultQuery("limit", "3")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid limit parameter",
		})
		return
	}

	carpools, err := controller.carpoolsService.GetTopLikedCarpools(limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch carpools",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "카풀 게시글을 좋아요 순으로 조회했습니다.",
		"data": gin.H{
			"page":  1, // 고정값 (필요 시 페이지네이션 추가 가능)
			"limit": limit,
			"total": len(carpools),
			"posts": carpools,
		},
	})
}

// 카풀 게시글 생성
func (controller *CarpoolsController) CreateCarpoolPost(ctx *gin.Context) {
	var request types.CreateCarpoolPostRequestDTO
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if err := controller.carpoolsService.CreateCarpoolsPost(request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create post"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Created Post"})
}
