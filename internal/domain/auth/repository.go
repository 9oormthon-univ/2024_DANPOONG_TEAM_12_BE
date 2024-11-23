package auth

import (
	"encoding/json"
	"fmt"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/url"
	"os"
)

type AuthRepository struct {
	DB     *gorm.DB
	client *resty.Client
}

func SetAuthRepository(DB *gorm.DB) *AuthRepository {
	return &AuthRepository{
		client: resty.New(),
		DB:     DB,
	}
}

func (r *AuthRepository) GetAccessToken(code string) (*types.KakaoTokenResponse, error) {
	tokenURL := "https://kauth.kakao.com/oauth/token"
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", os.Getenv("KAKAO_CLIENT_ID"))
	data.Set("redirect_uri", os.Getenv("KAKAO_REDIRECT_URI"))
	data.Set("code", code)

	// HTTP POST 요청
	resp, err := r.client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBody(data.Encode()).
		SetResult(&types.KakaoTokenResponse{}).
		Post(tokenURL)

	if err != nil {
		log.Printf("Failed to request access token: %v", err)
		return nil, err
	}

	// 응답 확인
	tokenResponse := resp.Result().(*types.KakaoTokenResponse)
	if tokenResponse.AccessToken == "" {
		log.Println("AccessToken is empty in the response")
		return nil, fmt.Errorf("failed to fetch access token: %s", resp.String())
	}

	log.Printf("AccessToken response: %+v", tokenResponse)
	return tokenResponse, nil
}

func (r *AuthRepository) GetUserInfo(accessToken string) (*types.KakaoUserResponse, error) {
	req, err := http.NewRequest("GET", "https://kapi.kakao.com/v2/user/me", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo types.KakaoUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}
