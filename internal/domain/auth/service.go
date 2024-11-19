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
