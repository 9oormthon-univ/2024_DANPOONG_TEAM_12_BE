package types

import "time"

type CoursesService interface {
	InjectAIService(service AIService)
	InjectRegionService(service RegionsService)
	RecommendCourses(req *RecommendCourseReq) (*Course, error)
}

// Course 구조체는 코스 정보를 나타냅니다.
type Course struct {
	CourseID  int64      `gorm:"primaryKey;autoIncrement" json:"-"`                            // 코스 ID, 자동 증가 및 기본 키
	Name      string     `gorm:"size:100" json:"name"`                                         // 코스 이름, 최대 100자
	StartDate *time.Time `json:"-"`                                                            // 시작 날짜, 시간 포인터 타입
	EndDate   *time.Time `json:"end_date"`                                                     // 종료 날짜, 시간 포인터 타입
	TotalTime *int       `json:"total_time"`                                                   // 총 시간, 포인터 타입
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"-"`                                      // 생성 시간, 자동 생성
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"-"`                                      // 수정 시간, 자동 업데이트
	UserID    int64      `json:"-"`                                                            // 사용자 ID
	Plans     []Plan     `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE" json:"plans"` // 연결된 계획들, 코스 삭제 시 함께 삭제
}

// Plan 구조체는 코스에 연결된 계획 정보를 나타냅니다.
type Plan struct {
	PlanID    int64     `gorm:"primaryKey;autoIncrement" json:"-"`                           // 계획 ID, 자동 증가 및 기본 키
	DayNumber string    `json:"day_number"`                                                  // 일자 번호
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`                                     // 생성 시간, 자동 생성
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`                                     // 수정 시간, 자동 업데이트
	CourseID  int64     `json:"-"`                                                           // 코스 ID
	Places    []Place   `gorm:"foreignKey:PlanID;constraint:OnDelete:CASCADE" json:"places"` // 연결된 장소들, 계획 삭제 시 함께 삭제
}

// Place 구조체는 계획에 연결된 장소 정보를 나타냅니다.
type Place struct {
	PlaceID int64  `gorm:"primaryKey;autoIncrement" json:"-"` // 장소 ID, 자동 증가 및 기본 키
	Name    string `gorm:"size:100" json:"title"`             // 장소 이름, 최대 100자
	// Duration  *int       // 소요 시간, 포인터 타입
	Description string `json:"description"`               // 장소 설명
	Address     string `json:"address"`                   // 장소 주소
	StartTime   string `json:"start_time"`                // 시작 시간
	EndTime     string `json:"end_time"`                  // 종료 시간
	Type        string `json:"type"`                      // 장소 유형
	MapX        string `json:"mapx"`                      // X 좌표 (경도 값, 위치 정보)
	MapY        string `json:"mapy"`                      // Y 좌표 (위도 값, 위치 정보)
	ImageURL    string `gorm:"size:255" json:"image_url"` // 이미지 URL, 최대 255자
	PlanID      int64  `json:"-"`                         // 계획 ID
}

// ---------------------------------------------------------
type RecommendCourseReq struct {
	AreaCode  string   `json:"area_code"`  // 지역 코드
	Interests []string `json:"interests"`  // 관심사
	StartTime string   `json:"start_time"` // 여행 시작 시간 (예: "10:00")
	EndTime   string   `json:"end_time"`
}

type CourseRecommendationAIRes struct {
	Title       string `json:"title"`       // 장소 이름
	Description string `json:"description"` // 장소 설명
	Address     string `json:"address"`     // 장소 주소
	StartTime   string `json:"start_time"`  // 시작 시간 (예: "10:00")
	EndTime     string `json:"end_time"`
	Type        string `json:"type"`
	ContentID   string `json:"content_id"`
}
