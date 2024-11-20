package types

import "time"

type Application struct {
	ApplicationID int       `gorm:"primaryKey;autoIncrement"`
	UserID        int       `gorm:"not null"`
	TargetID      int       `gorm:"not null"`
	TargetType    string    `gorm:"type:enum('matching','carpool');not null"`
	Status        string    `gorm:"type:enum('pending','approved','rejected');not null"`
	AppliedAt     time.Time `gorm:"autoCreateTime"`
}
