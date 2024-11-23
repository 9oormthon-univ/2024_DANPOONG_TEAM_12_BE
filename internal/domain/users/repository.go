package users

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"gorm.io/gorm"
	"log"
)

type UsersRepository struct {
	DB *gorm.DB
}

func SetUsersRepository(DB *gorm.DB) *UsersRepository {
	u := &UsersRepository{
		DB: DB,
	}
	user := types.User{
		Username: "testuser",
		Nickname: "JohnDoe",
	}

	if err := DB.Create(&user).Error; err != nil {
		log.Fatalf("Error inserting sample user data: %v", err)
	}
	return u
}

func (u *UsersRepository) GetByID(userID int64) (*types.User, error) {
	var user types.User
	// Attempt to find the user in the database
	if err := u.DB.First(&user, "user_id = ?", userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Return nil if user is not found
		}
		return nil, err // Return other errors
	}
	return &user, nil // Return the found user
}
