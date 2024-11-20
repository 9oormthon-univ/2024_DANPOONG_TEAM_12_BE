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

type Plan struct {
	PlanID    int       `gorm:"primaryKey;autoIncrement"`
	DayNumber int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CourseID  int       `gorm:"not null"`
	Places    []Place   `gorm:"foreignKey:PlanID"`
}

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
