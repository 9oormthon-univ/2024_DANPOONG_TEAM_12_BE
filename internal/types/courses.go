package types

import "time"

type CoursesService interface {
}

// Course 구조체는 코스 정보를 나타냅니다.
type Course struct {
	CourseID  int64      `gorm:"primaryKey;autoIncrement"` // 코스 ID, 자동 증가 및 기본 키
	Name      string     `gorm:"size:100;not null"`        // 코스 이름, 최대 100자, 필수 입력
	StartDate *time.Time // 시작 날짜, 시간 포인터 타입
	EndDate   *time.Time // 종료 날짜, 시간 포인터 타입
	TotalTime *int       // 총 시간, 포인터 타입
	CreatedAt time.Time  `gorm:"autoCreateTime"`                                  // 생성 시간, 자동 생성
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`                                  // 수정 시간, 자동 업데이트
	UserID    int64      `gorm:"not null"`                                        // 사용자 ID, 필수 입력
	Plans     []Plan     `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE"` // 연결된 계획들, 코스 삭제 시 함께 삭제
}

// Plan 구조체는 코스에 연결된 계획 정보를 나타냅니다.
type Plan struct {
	PlanID    int64     `gorm:"primaryKey;autoIncrement"`                      // 계획 ID, 자동 증가 및 기본 키
	DayNumber int       `gorm:"not null"`                                      // 일자 번호, 필수 입력
	CreatedAt time.Time `gorm:"autoCreateTime"`                                // 생성 시간, 자동 생성
	UpdatedAt time.Time `gorm:"autoUpdateTime"`                                // 수정 시간, 자동 업데이트
	CourseID  int64     `gorm:"not null"`                                      // 코스 ID, 필수 입력
	Places    []Place   `gorm:"foreignKey:PlanID;constraint:OnDelete:CASCADE"` // 연결된 장소들, 계획 삭제 시 함께 삭제
}

// Place 구조체는 계획에 연결된 장소 정보를 나타냅니다.
type Place struct {
	PlaceID   int64      `gorm:"primaryKey;autoIncrement"` // 장소 ID, 자동 증가 및 기본 키
	Name      string     `gorm:"size:100;not null"`        // 장소 이름, 최대 100자, 필수 입력
	StartTime *time.Time // 시작 시간, 시간 포인터 타입
	EndTime   *time.Time // 종료 시간, 시간 포인터 타입
	Duration  *int       // 소요 시간, 포인터 타입
	ImageURL  string     `gorm:"size:255"`  // 이미지 URL, 최대 255자
	Details   string     `gorm:"type:text"` // 상세 설명, 텍스트 타입
	PlanID    int64      `gorm:"not null"`  // 계획 ID, 필수 입력
}
