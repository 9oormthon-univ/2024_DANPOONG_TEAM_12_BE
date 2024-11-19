package controller

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/bootstrap/service"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/ai"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/auth"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/carpools"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/courses"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/matching"
	regions "github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/region"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/users"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controller struct {
	engine  *gin.Engine
	service *service.Service

	*users.UsersController
	*courses.CoursesController
	*ai.AIController
	*matching.MatchingController
	*carpools.CarpoolsController
	*regions.RegionsController
	*auth.AuthController
}

func SetController(engine *gin.Engine, service *service.Service) *Controller {
	c := &Controller{
		engine:  engine,
		service: service,
	}

	api := engine.Group("/api")

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	c.UsersController = users.SetUsersController(api, service.UsersService)
	c.CoursesController = courses.SetCoursesController(api, service.AIService)
	c.AIController = ai.SetAIController(api, service.AIService)
	c.MatchingController = matching.SetMatchingController(api, service.MatchingService)
	c.CarpoolsController = carpools.SetCarpoolsController(api, service.CarpoolsService)
	c.RegionsController = regions.SetRegionsController(api, service.RegionsService)
	c.AuthController = auth.SetAuthController(api, service.AuthService)

	return c
}
