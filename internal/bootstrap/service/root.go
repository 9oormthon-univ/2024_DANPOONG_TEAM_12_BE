package service

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/bootstrap/repository"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/ai"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/auth"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/carpools"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/courses"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/matching"
	regions "github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/region"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/users"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
)

type Service struct {
	types.UsersService
	types.AIService
	types.AuthService
	types.CarpoolsService
	types.MatchingService
	types.RegionsService
	types.CoursesService
}

func SetService(repository *repository.Repository) *Service {

	usersService := users.SetUsersService(repository.UsersRepository)
	aiService := ai.SetAIService(repository.AIRepository)
	authService := auth.SetAuthService(repository.AuthRepository)
	carpoolsService := carpools.SetCarpoolsService(repository.CarpoolsRepository)
	matchingService := matching.SetMatchingService(repository.MatchingRepository)
	regionsService := regions.SetRegionsService(repository.RegionsRepository)
	coursesService := courses.SetCoursesService(repository.CoursesRepository)

	s := &Service{
		AuthService:     authService,
		CarpoolsService: carpoolsService,
		MatchingService: matchingService,
		RegionsService:  regionsService,
		CoursesService:  coursesService,
		UsersService:    usersService,
		AIService:       aiService,
	}
	return s
}
