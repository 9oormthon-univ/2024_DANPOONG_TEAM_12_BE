package types

type UsersService interface {
}
type User struct {
	UserID        int            `gorm:"primaryKey;autoIncrement"`
	Username      string         `gorm:"size:50;not null"`
	Nickname      string         `gorm:"size:50;not null"`
	UserDetails   string         `gorm:"type:text"`
	UserInterests []UserInterest `gorm:"foreignKey:UserID"`
	Matchings     []Matching     `gorm:"foreignKey:AuthorID"`
	Carpools      []Carpool      `gorm:"foreignKey:AuthorID"`
}
