package ai

import "gorm.io/gorm"

type AIRepository struct {
	DB *gorm.DB
}

func SetAIRepository(DB *gorm.DB) *AIRepository {
	r := &AIRepository{
		DB: DB,
	}
	return r
}
