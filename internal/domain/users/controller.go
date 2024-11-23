package users

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersController struct {
	usersService types.UsersService
}

func SetUsersController(api *gin.RouterGroup, service types.UsersService) *UsersController {
	userController := &UsersController{
		usersService: service,
	}
	// 핸들러 등록
	api.POST("/users/signup", userController.SignUp)
	api.POST("/users/login", userController.Login)
	return userController
}

// 회원가입 API
func (uc *UsersController) SignUp(c *gin.Context) {
	var userRequest types.SignUpRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user, err := uc.usersService.SignUp(userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user})
}

// 로그인 API
func (uc *UsersController) Login(c *gin.Context) {
	var loginRequest types.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user, err := uc.usersService.Login(loginRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}
