package contents

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ContentsController struct {
	contentsService types.ContentsService
}

func SetContentsController(api *gin.RouterGroup, service types.ContentsService) *ContentsController {
	return &ContentsController{
		contentsService: service,
	}
}

func (c *ContentsController) GetAllContents(ctx *gin.Context) {
	contents, err := c.contentsService.GetAllContents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errer": "콘텐츠 조회 실패"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"contents": contents})
}

func (c *ContentsController) GetContentsByUD(ctx *gin.Context) {
	idParam := ctx.Param("id")
	contentID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 콘텐츠 ID"})
		return
	}

	content, err := c.contentsService.GetContentById(contentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "콘텐츠 조회 실패"})
		return
	}
	if content == nil {
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "콘텐츠를 찾을 수 없습니다."})

	ctx.JSON(http.StatusOK, content)
}
