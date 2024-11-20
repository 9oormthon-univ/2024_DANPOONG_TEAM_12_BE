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
