package types

import "time"

type CarpoolsService interface {
}

type Carpool struct {
	CarpoolID      int       `gorm:"primaryKey;autoIncrement"`
	Title          string    `gorm:"size:100;not null"`
	Details        string    `gorm:"type:text"`
	AuthorNickname string    `gorm:"size:50;not null"`
	AuthorID       int       `gorm:"not null"`
	Destination    string    `gorm:"size:100;not null"`
	Date           time.Time `gorm:"type:date;not null"`
	StartTime      *time.Time
	EndTime        *time.Time
	CreatedAt      time.Time     `gorm:"autoCreateTime"`
	UpdatedAt      time.Time     `gorm:"autoUpdateTime"`
	Status         string        `gorm:"type:enum('active','inactive');not null"`
	Likes          []CarpoolLike `gorm:"foreignKey:CarpoolID"`
}
