package types

import "time"

type MatchingService interface {
}

type Matching struct {
	MatchingID     int       `gorm:"primaryKey;autoIncrement"`
	Title          string    `gorm:"size:100;not null"`
	ImageURL       string    `gorm:"size:255"`
	Details        string    `gorm:"type:text"`
	AuthorNickname string    `gorm:"size:50;not null"`
	AuthorID       int       `gorm:"not null"`
	Destination    string    `gorm:"size:100;not null"`
	Date           time.Time `gorm:"type:date;not null"`
	StartTime      *time.Time
	EndTime        *time.Time
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
	Status         string         `gorm:"type:enum('active','inactive');not null"`
	Categories     []Category     `gorm:"foreignKey:MatchingID"`
	Likes          []MatchingLike `gorm:"foreignKey:MatchingID"`
}
