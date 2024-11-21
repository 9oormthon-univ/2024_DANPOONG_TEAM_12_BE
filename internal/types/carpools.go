package types

import "time"

type CarpoolsService interface {
	GetTopLikedCarpools(limit int) ([]CarpoolTopLikesResponseDTO, error)
}

// Carpool 구조체는 카풀 게시글 정보를 나타냅니다.
type Carpool struct {
	CarpoolID     int64                `gorm:"column:carpool_id;primaryKey;autoIncrement" json:"carpool_id"`  // 카풀 ID, 자동 증가 및 기본 키
	Title         string               `gorm:"column:title;size:100;not null" json:"title"`                   // 카풀 제목, 최대 100자, 필수 입력
	Details       string               `gorm:"column:details;type:text" json:"details"`                       // 상세 설명, 텍스트 타입
	StartLocation string               `gorm:"column:start_location;size:100;not null" json:"start_location"` // 출발지, 최대 100자, 필수 입력
	EndLocation   string               `gorm:"column:end_location;size:100;not null" json:"end_location"`     // 도착지, 최대 100자, 필수 입력
	StartTime     time.Time            `gorm:"column:start_time;type:datetime;not null" json:"start_time"`    // 출발 시간, 필수 입력
	Status        string               `gorm:"column:status;type:enum('active','completed','cancelled');not null" json:"status"`
	CreatedAt     time.Time            `gorm:"column:created_at;autoCreateTime" json:"created_at"`                    // 생성 시간, 자동 생성
	UpdatedAt     time.Time            `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`                    // 수정 시간, 자동 업데이트
	Applications  []CarpoolApplication `gorm:"foreignKey:CarpoolID;constraint:OnDelete:CASCADE;" json:"applications"` // 카풀 지원서들, 카풀 삭제 시 함께 삭제
	LikesModel    []CarpoolLike        `gorm:"foreignKey:CarpoolID;constraint:OnDelete:CASCADE;" json:"likes_model"`  // 카풀 좋아요들, 카풀 삭제 시 함께 삭제
	UserID        int64                `gorm:"column:user_id;not null" json:"user_id"`                                // 작성자 ID, 필수 입력
	User          User                 `gorm:"foreignKey:UserID" json:"user"`                                         // 작성자와의 관계
	Likes         int                  `gorm:"column:likes;default:0" json:"likes"`                                   // 좋아요 수, 기본값 0
}

// CarpoolApplication 구조체는 카풀 지원서 정보를 나타냅니다.
type CarpoolApplication struct {
	ApplicationID int64     `gorm:"primaryKey;autoIncrement"`                                              // 지원서 ID, 자동 증가 및 기본 키
	UserID        int64     `gorm:"not null"`                                                              // 지원자 ID, 필수 입력
	CarpoolID     int64     `gorm:"not null"`                                                              // 지원한 카풀 게시글 ID, 필수 입력
	AppliedAt     time.Time `gorm:"autoCreateTime"`                                                        // 지원 시간, 자동 생성
	User          User      `gorm:"foreignKey:UserID;references:UserID"`                                   // 지원자와의 관계
	Carpool       Carpool   `gorm:"foreignKey:CarpoolID;references:CarpoolID;constraint:OnDelete:CASCADE"` // 카풀 게시글과의 관계, 카풀 삭제 시 지원서도 삭제
}

// CarpoolLike 구조체는 카풀에 대한 좋아요 정보를 나타냅니다.
type CarpoolLike struct {
	LikeID    int64     `gorm:"primaryKey;autoIncrement"`                          // 좋아요 ID, 자동 증가 및 기본 키
	UserID    int64     `gorm:"not null"`                                          // 유저 ID, 필수 입력
	CarpoolID int64     `gorm:"not null"`                                          // 카풀 게시글 ID, 필수 입력
	LikedAt   time.Time `gorm:"autoCreateTime"`                                    // 좋아요 생성 시간, 자동 생성
	User      User      `gorm:"foreignKey:UserID";references:UserID`               // 유저와의 관계
	Carpool   Carpool   `gorm:"foreignKey:CarpoolID;constraint:OnDelete:CASCADE;"` // 카풀 게시글과의 관계, 카풀 삭제 시 좋아요도 삭제
}

type CarpoolTopLikesResponseDTO struct {
	CarpoolID     int64     `json:"carpool_id"`
	Title         string    `json:"title"`
	Details       string    `json:"details"`
	StartLocation string    `json:"start_location"`
	EndLocation   string    `json:"end_location"`
	StartTime     time.Time `json:"start_time"`
	LikesCount    int       `json:"likes_count"`
}
