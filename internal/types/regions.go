package types

import "time"

type RegionsService interface {
}

// LocalInfo 구조체는 지역 정보를 나타냅니다.
type LocalInfo struct {
	LocalInfoID  int64  `gorm:"primaryKey;autoIncrement"` // 지역 정보 ID, 자동 증가 및 기본 키
	Name         string `gorm:"size:100;not null"`        // 지역 이름, 최대 100자, 필수 입력
	LocalFeature string `gorm:"type:text"`                // 지역 특징, 텍스트 타입
	Specialty    string `gorm:"type:text"`                // 지역 특산물, 텍스트 타입
	Symbol       string `gorm:"type:text"`                // 지역 상징, 텍스트 타입
	Attraction   string `gorm:"type:text"`                // 지역 명소, 텍스트 타입
	Activity     string `gorm:"type:text"`                // 지역 활동, 텍스트 타입
}

// Content 구조체는 콘텐츠 정보를 나타냅니다.
type Content struct {
	ContentID   int64         `gorm:"primaryKey;autoIncrement"`                            // 콘텐츠 ID, 자동 증가 및 기본 키
	Name        string        `gorm:"size:100;not null"`                                   // 콘텐츠 이름, 최대 100자, 필수 입력
	Type        string        `gorm:"type:enum('festival','event','attraction');not null"` // 콘텐츠 유형, 필수 입력
	Details     string        `gorm:"type:text"`                                           // 콘텐츠 세부 정보, 텍스트 타입
	StartDate   time.Time     `gorm:"type:date;not null"`                                  // 시작 날짜, 필수 입력
	EndDate     time.Time     `gorm:"type:date;not null"`                                  // 종료 날짜, 필수 입력
	Location    string        `gorm:"size:100"`                                            // 위치, 최대 100자
	CreatedAt   time.Time     `gorm:"autoCreateTime"`                                      // 생성 시간, 자동 생성
	UpdatedAt   time.Time     `gorm:"autoUpdateTime"`                                      // 수정 시간, 자동 업데이트
	LocalInfoID int64         `gorm:"not null"`                                            // 연결된 지역 정보 ID, 필수 입력
	Likes       []ContentLike `gorm:"foreignKey:ContentID;constraint:OnDelete:CASCADE"`    // 연결된 좋아요, 콘텐츠 삭제 시 함께 삭제
}

// ContentLike 구조체는 콘텐츠에 대한 좋아요 정보를 나타냅니다.
type ContentLike struct {
	LikeID    int64     `gorm:"primaryKey;autoIncrement"`                          // 좋아요 ID, 자동 증가 및 기본 키
	UserID    int64     `gorm:"not null"`                                          // 사용자 ID, 필수 입력
	ContentID int64     `gorm:"not null"`                                          // 콘텐츠 ID, 필수 입력
	LikedAt   time.Time `gorm:"autoCreateTime"`                                    // 좋아요 생성 시간, 자동 생성
	User      User      `gorm:"foreignKey:UserID"`                                 // 사용자와의 관계
	Content   Content   `gorm:"foreignKey:ContentID;constraint:OnDelete:CASCADE;"` // 콘텐츠와의 관계, 콘텐츠 삭제 시 좋아요도 삭제
}
