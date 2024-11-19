package carpools

import "gorm.io/gorm"

type CarpoolsRepository struct {
	DB *gorm.DB
}

func SetCarpoolsRepository(DB *gorm.DB) *CarpoolsRepository {
	r := &CarpoolsRepository{
		DB: DB,
	}
	return r
}
