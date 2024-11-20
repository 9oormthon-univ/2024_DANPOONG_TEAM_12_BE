package types

import "time"

type MatchingService interface{}

// Matching 구조체는 매칭 정보를 나타냅니다.
type Matching struct {
	MatchingID   int64                 `gorm:"primaryKey;autoIncrement"`
	Title        string                `gorm:"size:100;not null"`
	ImageURL     string                `gorm:"size:255"`
	Details      string                `gorm:"type:text"`
	UserNickname string                `gorm:"size:50;not null"`
	UserID       int64                 `gorm:"not null"`
	Destination  string                `gorm:"size:100;not null"`
	Date         time.Time             `gorm:"type:date;not null"`
	StartTime    *time.Time            `gorm:"type:time"`
	EndTime      *time.Time            `gorm:"type:time"`
	CreatedAt    time.Time             `gorm:"autoCreateTime"`
	UpdatedAt    time.Time             `gorm:"autoUpdateTime"`
	User         User                  `gorm:"foreignKey:UserID;references:UserID"`
	Categories   []Category            `gorm:"foreignKey:MatchingID;constraint:OnDelete:CASCADE;"`
	LikesModel   []MatchingLike        `gorm:"foreignKey:MatchingID;constraint:OnDelete:CASCADE;"`
	Applications []MatchingApplication `gorm:"foreignKey:MatchingID;constraint:OnDelete:CASCADE;"`
	Likes        int                   `gorm:"default:0"`
}

type MatchingApplication struct {
	ApplicationID int64     `gorm:"primaryKey;autoIncrement"`
	UserID        int64     `gorm:"not null"`
	MatchingID    int64     `gorm:"not null"`
	AppliedAt     time.Time `gorm:"autoCreateTime"`
	User          User      `gorm:"foreignKey:UserID"`
	Matching      Matching  `gorm:"constraint:OnDelete:CASCADE;"`
}

type Category struct {
	CategoryID int64    `gorm:"primaryKey;autoIncrement"`
	Name       string   `gorm:"size:50;not null"`
	MatchingID int64    `gorm:"not null"`
	Matching   Matching `gorm:"constraint:OnDelete:CASCADE;"`
}

type MatchingLike struct {
	LikeID     int64     `gorm:"primaryKey;autoIncrement"`
	UserID     int64     `gorm:"not null"`
	MatchingID int64     `gorm:"not null"`
	LikedAt    time.Time `gorm:"autoCreateTime"`
	User       User      `gorm:"foreignKey:UserID"`
	Matching   Matching  `gorm:"constraint:OnDelete:CASCADE;"`
}
