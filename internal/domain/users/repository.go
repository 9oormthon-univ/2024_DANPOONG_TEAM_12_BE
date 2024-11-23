package users

import (
	"encoding/json"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"gorm.io/gorm"
)

type UsersRepository struct {
	DB *gorm.DB
}

func SetUsersRepository(DB *gorm.DB) *UsersRepository {
	u := &UsersRepository{
		DB: DB,
	}
	return u
}

// GetByID - 사용자 ID로 사용자 정보를 조회
func (r *UsersRepository) GetByID(userID int64) (*types.User, error) {
	var user types.User
	err := r.DB.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 카카오 ID로 사용자 조회
func (r *UsersRepository) FindByKakaoID(kakaoID int64) (*types.User, error) {
	var user types.User
	if err := r.DB.Where("kakao_id = ?", kakaoID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // 사용자가 없으면 nil 반환
		}
		return nil, err
	}
	return &user, nil
}

// 사용자 생성
func (r *UsersRepository) Create(user *types.User) error {
	return r.DB.Create(user).Error
}

// 사용자 조회
func (r *UsersRepository) FindByUsername(username string) (*types.User, error) {
	var user types.User
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UsersRepository) AddUserInterests(userID int64, interests []string) error {
	// 사용자 엔티티를 찾음
	var user types.User
	if err := r.DB.First(&user, userID).Error; err != nil {
		return err
	}

	// `Interests`를 갱신
	user.Interests = interests

	// JSON으로 저장
	interestsJSON, err := json.Marshal(interests)
	if err != nil {
		return err
	}

	// interests를 JSON 형태로 업데이트
	if err := r.DB.Model(&user).Update("user_details", string(interestsJSON)).Error; err != nil {
		return err
	}

	return nil
}
