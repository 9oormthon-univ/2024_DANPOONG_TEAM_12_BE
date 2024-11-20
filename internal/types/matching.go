package types

import "time"

type MatchingService interface {
}

type Matching struct {
	MatchingID     int64          `gorm:"primaryKey;autoIncrement"` // 매칭 ID
	Title          string         `gorm:"size:100;not null"`        // 매칭 제목
	ImageURL       string         `gorm:"size:255"`                 // 이미지 URL
	Details        string         `gorm:"type:text"`                // 상세 설명
	AuthorNickname string         `gorm:"size:50;not null"`         // 작성자 닉네임
	AuthorID       int64          `gorm:"not null"`                 // 작성자 ID
	Destination    string         `gorm:"size:100;not null"`        // 여행지
	Date           time.Time      `gorm:"type:date;not null"`       // 여행 날짜
	StartTime      *time.Time     `gorm:"type:time"`                // 시작 시간
	EndTime        *time.Time     `gorm:"type:time"`                // 종료 시간
	CreatedAt      time.Time      `gorm:"autoCreateTime"`           // 생성 시간
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`           // 수정 시간
	Status         string         `gorm:"size:20;not null"`         // 상태 (active/inactive)
	Categories     []Category     `gorm:"foreignKey:MatchingID"`    // 매칭과 연결된 카테고리들
	Likes          []MatchingLike `gorm:"foreignKey:MatchingID"`    // 매칭과 연결된 좋아요들
}

type Category struct {
	CategoryID int64    `gorm:"primaryKey;autoIncrement"`     // 카테고리 ID
	Name       string   `gorm:"size:50;not null"`             // 카테고리 이름
	MatchingID int64    `gorm:"not null"`                     // 매칭 ID
	Matching   Matching `gorm:"constraint:OnDelete:CASCADE;"` // 매칭과의 관계 (매칭 삭제 시 카테고리도 삭제)
}

type MatchingLike struct {
	LikeID     int64     `gorm:"primaryKey;autoIncrement"`     // 좋아요 ID
	UserID     int64     `gorm:"not null"`                     // 사용자 ID
	MatchingID int64     `gorm:"not null"`                     // 매칭 ID
	LikedAt    time.Time `gorm:"autoCreateTime"`               // 좋아요 생성 시간
	Matching   Matching  `gorm:"constraint:OnDelete:CASCADE;"` // 매칭과의 관계 (매칭 삭제 시 좋아요도 삭제)
}
