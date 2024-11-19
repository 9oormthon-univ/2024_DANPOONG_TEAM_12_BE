package types

import "time"

type MatchingLike struct {
	LikeID     int       `gorm:"primaryKey;autoIncrement"`
	UserID     int       `gorm:"not null"`
	MatchingID int       `gorm:"not null"`
	LikedAt    time.Time `gorm:"autoCreateTime"`
}

type CarpoolLike struct {
	LikeID    int       `gorm:"primaryKey;autoIncrement"`
	UserID    int       `gorm:"not null"`
	CarpoolID int       `gorm:"not null"`
	LikedAt   time.Time `gorm:"autoCreateTime"`
}
