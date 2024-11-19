package types

import "time"

type Plan struct {
	PlanID    int       `gorm:"primaryKey;autoIncrement"`
	DayNumber int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CourseID  int       `gorm:"not null"`
	Places    []Place   `gorm:"foreignKey:PlanID"`
}
