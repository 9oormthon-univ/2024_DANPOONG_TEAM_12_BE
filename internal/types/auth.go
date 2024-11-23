package types

type AuthService interface {
	HandleKakaoLogin(code string) (*KakaoUserResponse, error)
	GetAccessToken(code string) (*KakaoTokenResponse, error)
	GetKakaoUserInfo(accessToken string) (*KakaoUserResponse, error)
}

type KakaoUserInfo struct {
	ID           int64  `json:"id"`
	Nickname     string `json:"properties.nickname"`
	Email        string `json:"kakao_account.email"`
	ProfileImage string `json:"properties.profile_image"`
}

type KakaoTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type KakaoUserResponse struct {
	ID           int64  `json:"id"`
	Nickname     string `json:"properties.nickname"`
	Email        string `json:"kakao_account.email"`
	ProfileImage string `json:"properties.profile_image"`
}
