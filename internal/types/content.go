package types

import "time"

type Content struct {
	ContentID   int       `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"size:100;not null"`
	Type        string    `gorm:"type:enum('festival','event','attraction');not null"`
	Details     string    `gorm:"type:text"`
	StartDate   time.Time `gorm:"type:date;not null"`
	EndDate     time.Time `gorm:"type:date;not null"`
	Location    string    `gorm:"size:100"`
	Status      string    `gorm:"type:enum('active','inactive');not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	LocalInfoID int       `gorm:"not null"`
}
