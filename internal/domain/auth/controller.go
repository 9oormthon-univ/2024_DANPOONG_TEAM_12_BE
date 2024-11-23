package auth

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type AuthController struct {
	authService types.AuthService
	userService types.UsersService
}

func SetAuthController(api *gin.RouterGroup, service types.AuthService) *AuthController {
	authController := &AuthController{
		authService: service,
	}
	// 핸들러 등록
	api.GET("/auth/kakao/login", authController.KakaoLogin)
	api.GET("/auth/kakao/callback", authController.KakaoCallback)

	return authController
}

// 1. 카카오 로그인 요청
func (ac *AuthController) KakaoLogin(c *gin.Context) {
	clientID := os.Getenv("KAKAO_CLIENT_ID")
	redirectURI := os.Getenv("KAKAO_REDIRECT_URI")

	if clientID == "" || redirectURI == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing environment variables"})
		return
	}

	kakaoAuthURL := "https://kauth.kakao.com/oauth/authorize?response_type=code" +
		"&client_id=" + clientID +
		"&redirect_uri=" + redirectURI

	c.Redirect(http.StatusFound, kakaoAuthURL)
}

func (ac *AuthController) KakaoCallback(c *gin.Context) {
	log.Println("Starting KakaoCallback...")

	// Step 1: Authorization Code 받기
	code := c.Query("code")
	log.Println("Received code:", code)
	if code == "" {
		log.Println("Authorization code is missing")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization code not provided"})
		return
	}

	// Step 2: AccessToken 요청
	tokenResponse, err := ac.authService.GetAccessToken(code)
	if err != nil {
		log.Println("Failed to fetch access token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch access token",
			"details": err.Error(),
		})
		return
	}

	log.Println("AccessToken:", tokenResponse.AccessToken)

	// Step 3: 사용자 정보 가져오기
	kakaoUser, err := ac.authService.GetKakaoUserInfo(tokenResponse.AccessToken)
	if err != nil {
		log.Println("Failed to fetch user info:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch user info",
			"details": err.Error(),
		})
		return
	}

	log.Println("Kakao User Info:", kakaoUser)

	// Step 4: 회원가입 또는 로그인 처리
	//user, err := ac.userService.RegisterOrLogin(kakaoUser)
	//if err != nil {
	//	log.Println("Failed to register or login user:", err)
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"error":   "Failed to register or login user",
	//		"details": err.Error(),
	//	})
	//	return
	//}
	//
	//// Step 5: 응답 반환 (JWT 토큰 등 추가 가능)
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "Login successful",
	//	"user":    user,
	//})
}
