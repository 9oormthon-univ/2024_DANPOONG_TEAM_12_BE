package repository

import (
	"fmt"
	"log"

	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/config"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/db"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/ai"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/auth"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/carpools"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/courses"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/matching"
	regions "github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/region"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/domain/users"
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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

	err := db.ResetDatabase(r.DB)
	if r.DB == nil {
		return nil, fmt.Errorf("failed to connect to the database")
	}

	// AutoMigrate를 사용하여 테이블 생성
	err = initializeTables(r.DB)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("데이터베이스 테이블 초기화 실패 :%w", err)
	}

	log.Println("테이블 초기화 성공!!")

	// 리포지토리 설정
	r.UsersRepository = users.SetUsersRepository(r.DB)
	r.AIRepository = ai.SetAIRepository(r.DB)
	r.AuthRepository = auth.SetAuthRepository(r.DB)
	r.RegionsRepository = regions.SetRegionsRepository(r.DB)
	r.CarpoolsRepository = carpools.SetCarpoolsRepository(r.DB)
	r.CoursesRepository = courses.SetCoursesRepository(r.DB)
	r.MatchingRepository = matching.SetMatchingRepository(r.DB)

	// db.DbCRUDtest(r.DB)
	return r, nil
}

func (r *Repository) ConnectToDB(cfg *config.Config) *gorm.DB {
	// 데이터베이스 설정 확인
	if cfg.DBUser == "" || cfg.DBPassword == "" || cfg.DBHost == "" || cfg.DBPort == 0 || cfg.DBName == "" {
		log.Println("invalid database configuration")
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 단수형 테이블 이름 사용
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 외래 키 비활성화
	})

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

	log.Println("데이터베이스 연결 성공!!!!!!!1")

	return db
}

func initializeTables(db *gorm.DB) error {
	// 테이블 초기화
	err := db.AutoMigrate(
		&types.User{},
		&types.Interest{},
		&types.Matching{},
		&types.Category{},
		&types.Carpool{},
		&types.MatchingLike{},
		&types.CarpoolLike{},
		&types.Course{},
		&types.Plan{},
		&types.Place{},
		&types.LocalInfo{},
		&types.Content{},
		&types.ContentLike{},
		&types.MatchingApplication{},
		&types.CarpoolApplication{},
	)
	if err != nil {
		return err
	}

	// 외래 키 추가 (SQL 실행)
	err = addForeignKeys(db)
	if err != nil {
		return err
	}

	return nil
}

func addForeignKeys(db *gorm.DB) error {
	foreignKeyStatements := []string{
		`ALTER TABLE matching ADD CONSTRAINT fk_user_matching FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE;`,
		`ALTER TABLE matching_like ADD CONSTRAINT fk_user_matching_likes FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE;`,
		`ALTER TABLE matching_like ADD CONSTRAINT fk_matching_matching_likes FOREIGN KEY (matching_id) REFERENCES matching(matching_id) ON DELETE CASCADE;`,
		`ALTER TABLE matching_application ADD CONSTRAINT fk_user_matching_applications FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE;`,
		`ALTER TABLE matching_application ADD CONSTRAINT fk_matching_matching_applications FOREIGN KEY (matching_id) REFERENCES matching(matching_id) ON DELETE CASCADE;`,
		`ALTER TABLE carpool ADD CONSTRAINT fk_user_carpools FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE;`,
		`ALTER TABLE carpool_like ADD CONSTRAINT fk_user_carpool_likes FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE;`,
		`ALTER TABLE carpool_like ADD CONSTRAINT fk_carpool_carpool_likes FOREIGN KEY (carpool_id) REFERENCES carpool(carpool_id) ON DELETE CASCADE;`,
	}

	for _, stmt := range foreignKeyStatements {
		if err := db.Exec(stmt).Error; err != nil {
			log.Printf("외래 키 추가 실패: %v", err)
			return err
		}
	}

	log.Println("외래 키 추가 완료")
	return nil
}
