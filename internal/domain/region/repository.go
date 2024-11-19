package regions

import "gorm.io/gorm"

type RegionsRepository struct {
	DB *gorm.DB
}

func SetRegionsRepository(DB *gorm.DB) *RegionsRepository {
	r := &RegionsRepository{
		DB: DB,
	}
	return r
}
