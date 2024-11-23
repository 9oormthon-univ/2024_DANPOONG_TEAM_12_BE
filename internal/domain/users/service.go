package users

import (
	"errors"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"golang.org/x/crypto/bcrypt"
)

type usersService struct {
	usersRepository *UsersRepository
}

func SetUsersService(repository *UsersRepository) types.UsersService {
	u := &usersService{
		usersRepository: repository,
	}
	return u
}

// 회원가입
func (s *usersService) SignUp(request types.SignUpRequest) (*types.User, error) {
	// 비밀번호 해싱
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 사용자 정보 생성
	user := &types.User{
		Username:     request.Username,
		Nickname:     request.Nickname,
		Email:        request.Email,
		Password:     string(hashedPassword),
		UserDetails:  request.UserDetails,
		ProfileImage: request.ProfileImage,
		Interests:    request.Interests,
	}

	// 사용자 생성
	if err := s.usersRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// 로그인
func (s *usersService) Login(request types.LoginRequest) (*types.User, error) {
	// 사용자 조회
	user, err := s.usersRepository.FindByUsername(request.Username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	// 비밀번호 검증
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, errors.New("invalid password")
	}
	user.Interests = []string{"맛집 투어", "혼자서", "힐링"}
	return user, nil
}
