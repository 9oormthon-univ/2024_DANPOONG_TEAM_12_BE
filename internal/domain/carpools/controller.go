package carpools

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CarpoolsController struct {
	carpoolsService types.CarpoolsService
}

func SetCarpoolsController(api *gin.RouterGroup, service types.CarpoolsService) *CarpoolsController {
	c := &CarpoolsController{
		carpoolsService: service,
	}
	// 핸들러 등록
	api.GET("/carpools/posts/sorted-by-likes", c.GetTopLikedCarpools)

	return c
}

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
