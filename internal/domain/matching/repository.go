package matching

import "gorm.io/gorm"

type MatchingRepository struct {
	DB *gorm.DB
}

func SetMatchingRepository(DB *gorm.DB) *MatchingRepository {
	r := &MatchingRepository{
		DB: DB,
	}
	return r
}
