package repository

import (
	"fmt"
	"log"

	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/config"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/ai"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/auth"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/carpools"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/courses"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/matching"
	regions "github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/region"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	*users.UsersRepository
	*ai.AIRepository
	*auth.AuthRepository
	*regions.RegionsRepository
	*carpools.CarpoolsRepository
	*courses.CoursesRepository
	*matching.MatchingRepository
	DB *gorm.DB
}

func SetRepository(cfg *config.Config) (*Repository, error) {
	r := &Repository{}
	r.DB = r.ConnectToDB(cfg)
	r.UsersRepository = users.SetUsersRepository(r.DB)

	return r, nil
}

func (r *Repository) ConnectToDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// 연결 상태 확인 (optional)
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get DB: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	return db
}
