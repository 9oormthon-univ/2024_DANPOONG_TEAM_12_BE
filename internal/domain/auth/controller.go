package auth

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService types.AuthService
}

func SetAuthController(api *gin.RouterGroup, service types.AuthService) *AuthController {
	c := &AuthController{
		authService: service,
	}
	// 핸들러 등록
	return c
}
