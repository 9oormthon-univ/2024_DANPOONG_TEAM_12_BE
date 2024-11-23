package cmd

import (
	"fmt"

	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/config"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/bootstrap"
	"github.com/gin-gonic/gin"
)

// 서버 실행
func Run() {
	engine := gin.New()

	gin.SetMode(gin.DebugMode)

	// engine.Use(middleware.CORSMiddleware())
	cfg, err := config.SetConfig()
	if err != nil {
		panic(err)
	}

	if err := bootstrap.InitApplication(engine, cfg); err != nil {
		panic(err)
	}

	if err := engine.Run(fmt.Sprintf(":%s", cfg.AppPort)); err != nil {
		panic(err)
	}
}
