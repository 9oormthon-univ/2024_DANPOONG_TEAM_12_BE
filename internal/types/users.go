package types

type UsersService interface {
}

type User struct {
	UserID      int        `gorm:"primaryKey;autoIncrement"`
	Username    string     `gorm:"size:50;not null"`
	Nickname    string     `gorm:"size:50;not null"`
	UserDetails string     `gorm:"type:text"`
	Interests   []Interest `gorm:"many2many:users_interests"` // 다대다 관계
}

type Interest struct {
	InterestID int    `gorm:"primaryKey;autoIncrement"`
	Name       string `gorm:"size:50;not null"`
	Users      []User `gorm:"many2many:users_interests"` // 다대다 관계
}
