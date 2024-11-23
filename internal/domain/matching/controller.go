package matching

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MatchingController struct {
	matchingService types.MatchingService
}

func SetMatchingController(api *gin.RouterGroup, service types.MatchingService) *MatchingController {
	c := &MatchingController{
		matchingService: service,
	}
	// 핸들러 등록
	// 매칭 게시글 좋아요 순 조회
	api.GET("/matching/posts/sorted-by-likes", c.GetTopLikedMatchingPosts)
	return c
}

// 매칭 게시글 좋아요 순 조회
func (controller *MatchingController) GetTopLikedMatchingPosts(ctx *gin.Context) {
	limitStr := ctx.DefaultQuery("limit", "3") // 기본값 3
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid limit parameter"})
		return
	}

	// 서비스 호출
	posts, err := controller.matchingService.GetTopLikeMatchingPosts(limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to fetch posts"})
		return
	}

	// 성공 응답
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "매칭 게시글을 좋아요 순으로 조회했습니다.",
		"data": gin.H{
			"page":  1, // 고정값 (확장 가능)
			"limit": limit,
			"total": len(posts),
			"posts": posts,
		},
	})
}
