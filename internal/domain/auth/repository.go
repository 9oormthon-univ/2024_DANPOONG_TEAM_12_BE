package auth

import "gorm.io/gorm"

type AuthRepository struct {
	DB *gorm.DB
}

func SetAuthRepository(DB *gorm.DB) *AuthRepository {
	r := &AuthRepository{
		DB: DB,
	}
	return r
}
