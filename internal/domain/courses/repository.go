package courses

import "gorm.io/gorm"

type CoursesRepository struct {
	DB *gorm.DB
}

func SetCoursesRepository(DB *gorm.DB) *CoursesRepository {
	r := &CoursesRepository{
		DB: DB,
	}
	return r
}
