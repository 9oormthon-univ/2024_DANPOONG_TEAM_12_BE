package courses

import "github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"

type CoursesService struct {
	coursesRepository *CoursesRepository
}

func SetCoursesService(repository *CoursesRepository) types.CoursesService {
	r := &CoursesService{
		coursesRepository: repository,
	}
	return r
}
