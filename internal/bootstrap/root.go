package bootstrap

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/config"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/bootstrap/controller"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/bootstrap/repository"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/bootstrap/service"
	"github.com/gin-gonic/gin"
)

func InitApplication(engine *gin.Engine, cfg *config.Config) error {
	r, err := repository.SetRepository(cfg)
	if err != nil {
		return err
	}
	s := service.SetService(r)
	controller.SetController(engine, s)

	return nil
}
