package db

import (
	"fmt"
	"log"

	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"

	"gorm.io/gorm"
)

func ResetDatabase(db *gorm.DB) error {
	log.Println("데이터베이스 초기화 시작...")

	// 외래 키 제약 조건 비활성화
	if err := db.Exec("SET FOREIGN_KEY_CHECKS = 0;").Error; err != nil {
		return fmt.Errorf("외래 키 제약 조건 비활성화 실패: %w", err)
	}

	// 테이블 삭제 순서: 종속성 없는 테이블부터 시작
	models := []interface{}{
		&types.MatchingLike{},
		&types.CarpoolLike{},
		&types.MatchingApplication{},
		&types.CarpoolApplication{},
		&types.Category{},
		&types.Place{},
		&types.Plan{},
		&types.Course{},
		&types.Matching{},
		&types.Carpool{},
		&types.Content{},
		&types.LocalInfo{},
		&types.Interest{},
		&types.User{},
	}

	for _, model := range models {
		log.Printf("테이블 삭제 시도: %T\n", model)
		if err := db.Migrator().DropTable(model); err != nil {
			log.Printf("테이블 삭제 실패 (%T): %v\n", model, err)
		} else {
			log.Printf("테이블 삭제 성공 (%T)\n", model)
		}
	}

	// 외래 키 제약 조건 다시 활성화
	if err := db.Exec("SET FOREIGN_KEY_CHECKS = 1;").Error; err != nil {
		return fmt.Errorf("외래 키 제약 조건 활성화 실패: %w", err)
	}

	log.Println("데이터베이스 전체 테이블 삭제 완료!")
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
