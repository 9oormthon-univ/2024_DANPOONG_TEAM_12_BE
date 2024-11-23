package matching

import (
	"net/http"
	"strconv"

	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/util"
	"github.com/gin-gonic/gin"
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
	// 매칭 게시글 작성
	api.POST("/matching/posts", c.CreateMatchingPost)
	// 내가 작성한 매칭 게시글 목록 조회
	api.GET("/matching/posts/me", c.GetUserMatchingPosts)
	// 매칭 지원하기
	api.POST("/matching/:matchingID/applications", c.CreateMatchingApplication)
	// AI 기반 매칭 게시글 추천
	api.GET("/matching/posts/ai", c.GetRecommendMatchingPost)
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

// 매칭 게시글 생성
func (controller *MatchingController) CreateMatchingPost(ctx *gin.Context) {
	var request types.CreateMatchingPostRequestDTO
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	if err := controller.matchingService.CreateMatchingPost(request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create post"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Created post"})
}

// 작성한 매칭 게시글 목록 조회
func (controller *MatchingController) GetUserMatchingPosts(ctx *gin.Context) {
	var requestDTO types.GetUserMatchingPostRequestDTO
	if err := ctx.ShouldBindJSON(&requestDTO); err != nil {
		// 요청 바디가 유효하지 않으면 에러 응답
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	// 서비스 계층 호출
	posts, err := controller.matchingService.GetUserMatchingPosts(requestDTO.UserID)
	if err != nil {
		// 서비스 계층에서 에러 발생 시 응답
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 매칭 게시글이 없는 경우 처리
	if len(posts) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No matching posts found for the user."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"matching_posts": posts})
}

// CreateMatchingApplication - 매칭 지원서 생성
func (controller *MatchingController) CreateMatchingApplication(ctx *gin.Context) {
	var requestDTO types.MatchingApplicationRequestDTO

	if err := ctx.ShouldBindJSON(&requestDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body", "error": err.Error()})
		return
	}

	// 매칭 ID 가져오기
	matchingID, err := strconv.ParseInt(ctx.Param("matchingID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid matching ID"})
		return
	}

	// userID 유효성 검사
	if requestDTO.UserID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "User ID is invalid or missing in the request body."})
		return
	}

	// 서비스 호출 - user_id 사용
	application, err := controller.matchingService.CreateMatchingApplication(requestDTO, matchingID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	// 성공 응답
	ctx.JSON(http.StatusCreated, application)
}

func (controller *MatchingController) GetRecommendMatchingPost(ctx *gin.Context) {
	// 로그인 기능 되면 사용자 관심사 불러오기
	interest := []string{"등산", "친구"}
	location := "서울"
	// 위치 자동 추적 기능 완성 되면 위치가져오기
	posts, err := controller.matchingService.RecommendMatchingPosts(location, interest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.SuccessResponse("매칭 게시글 추천 성공", posts))
	}
}
