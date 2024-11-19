package users

import "gorm.io/gorm"

type UsersRepository struct {
	DB *gorm.DB
}

func SetUsersRepository(DB *gorm.DB) *UsersRepository {
	u := &UsersRepository{
		DB: DB,
	}
	return u
}
