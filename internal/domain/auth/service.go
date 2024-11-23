package auth

import "github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"

type AuthService struct {
	authRepository *AuthRepository
}

func SetAuthService(repository *AuthRepository) types.AuthService {
	r := &AuthService{
		authRepository: repository,
	}
	return r
}
func (s *AuthService) GetAccessToken(code string) (*types.KakaoTokenResponse, error) {
	resp, err := s.authRepository.GetAccessToken(code)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) GetKakaoUserInfo(accessToken string) (*types.KakaoUserResponse, error) {
	resp, err := s.authRepository.GetUserInfo(accessToken)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) HandleKakaoLogin(code string) (*types.KakaoUserResponse, error) {
	// Step 1: Get access token
	tokenResponse, err := s.authRepository.GetAccessToken(code)
	if err != nil {
		return nil, err
	}

	// Step 2: Get user info
	userInfo, err := s.authRepository.GetUserInfo(tokenResponse.AccessToken)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
