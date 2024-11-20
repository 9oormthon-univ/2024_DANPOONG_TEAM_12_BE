package db

import (
	"fmt"
	"log"

	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"

	"gorm.io/gorm"
)

func ResetDatabase(db *gorm.DB) error {
	// 모델 기반으로 삭제할 테이블
	models := []interface{}{
		&types.User{},
		&types.Matching{},
		&types.Category{},
		&types.MatchingLike{},
		&types.Carpool{},
		&types.CarpoolLike{},
		&types.Course{},
		&types.Plan{},
		&types.Place{},
		&types.LocalInfo{},
		&types.Content{},
		&types.Application{},
		&types.Interest{},
	}

	if err := db.Migrator().DropTable(models...); err != nil {
		return fmt.Errorf("모델 기반 테이블 삭제 실패: %w", err)
	}
	log.Println("데이터베이스 테이블 삭제 성공!!")

	return nil
}

func DbCRUDtest(db *gorm.DB) error {
	// Create
	user := &types.User{
		Username:    "johndoe",
		Nickname:    "John",
		UserDetails: "This is a test user.",
	}
	if err := db.Create(user).Error; err != nil {
		return fmt.Errorf("Create 실패: %w", err)
	}
	fmt.Printf("Create 성공: %+v\n", user)

	// Read
	var readUser types.User
	if err := db.First(&readUser, user.UserID).Error; err != nil {
		return fmt.Errorf("Read 실패: %w", err)
	}
	fmt.Printf("Read 성공: %+v\n", readUser)

	// Update
	readUser.Nickname = "Johnny"
	if err := db.Save(&readUser).Error; err != nil {
		return fmt.Errorf("Update 실패: %w", err)
	}
	fmt.Printf("Update 성공: %+v\n", readUser)

	// Delete
	if err := db.Delete(&types.User{}, readUser.UserID).Error; err != nil {
		return fmt.Errorf("Delete 실패: %w", err)
	}
	fmt.Println("Delete 성공")

	// Verify Delete
	var verifyUser types.User
	if err := db.First(&verifyUser, readUser.UserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Verify Delete 성공: 레코드가 삭제되었습니다.")
			return nil
		}
		return fmt.Errorf("Verify Delete 실패: %w", err)
	}

	return fmt.Errorf("Verify Delete 실패: 레코드가 여전히 존재합니다")
}
