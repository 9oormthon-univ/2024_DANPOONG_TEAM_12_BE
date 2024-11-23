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

	api.GET("/regions//areaBasedList", r.GetareaBasedList)
	api.GET("/regions/test/detailCommon", r.GetDetailCommon)
	return r
}

func (r *RegionsController) GetareaBasedList(ctx *gin.Context) {
	var req types.AreaBasedListRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	result, err := r.regionsService.GetAreaBasedList(types.AreaCode(req.AreaCode), types.ContentType(req.ContentType))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result[0].ContentID,
	})
}

func (r *RegionsController) GetDetailCommon(ctx *gin.Context) {

	result, err := r.regionsService.GetDetailCommon("3417988")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "errofr",
		},
		)
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "success",
		"data":    result.Title,
	},
	)
}
