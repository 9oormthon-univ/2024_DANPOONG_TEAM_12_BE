package types

import "time"

type Place struct {
	PlaceID   int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:100;not null"`
	StartTime *time.Time
	EndTime   *time.Time
	Duration  *int
	ImageURL  string `gorm:"size:255"`
	Details   string `gorm:"type:text"`
	PlanID    int    `gorm:"not null"`
}
