package users

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	usersService types.UsersService
}

func SetUsersController(api *gin.RouterGroup, service types.UsersService) *UsersController {
	u := &UsersController{
		usersService: service,
	}
	// 핸들러 등록

	return u
}
