package types

import "time"

type CoursesService interface {
}

type Course struct {
	CourseID  int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:100;not null"`
	StartDate *time.Time
	EndDate   *time.Time
	TotalTime *int
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	UserID    int       `gorm:"not null"`
	Plans     []Plan    `gorm:"foreignKey:CourseID"`
}
