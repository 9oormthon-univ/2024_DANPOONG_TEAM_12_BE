package types

type Interest struct {
	InterestID    int            `gorm:"primaryKey;autoIncrement"`
	Name          string         `gorm:"size:50;not null"`
	UserInterests []UserInterest `gorm:"foreignKey:InterestID"`
}

type UserInterest struct {
	UserID     int      `gorm:"primaryKey"`
	InterestID int      `gorm:"primaryKey"`
	User       User     `gorm:"foreignKey:UserID"`
	Interest   Interest `gorm:"foreignKey:InterestID"`
}
