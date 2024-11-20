package regions

import (
	"net/http"

	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/gin-gonic/gin"
)

type RegionsController struct {
	regionsService types.RegionsService
}

func SetRegionsController(api *gin.RouterGroup, service types.RegionsService) *RegionsController {
	r := &RegionsController{
		regionsService: service,
	}
	// 핸들러 등록

	api.GET("/regions/test", r.GetareaBasedList)
	return r
}

func (r *RegionsController) GetareaBasedList(ctx *gin.Context) {
	result, err := r.regionsService.GetareaBasedList(types.Busan)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		},
		)
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "success",
		"data":    result[0].ContentID,
	},
	)
}
