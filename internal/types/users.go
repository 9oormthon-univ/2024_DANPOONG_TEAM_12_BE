// types 패키지 정의
package types

// UsersService 인터페이스 정의
type UsersService interface {
}

type User struct {
	UserID               int64                 `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username             string                `gorm:"size:50;not null" json:"username"`
	Nickname             string                `gorm:"size:50;not null" json:"nickname"`
	UserDetails          string                `gorm:"type:text" json:"user_details"`
	Interests            []Interest            `gorm:"many2many:user_interest;constraint:OnDelete:CASCADE;" json:"interests"`
	Matchings            []Matching            `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE;" json:"matchings"`
	MatchingApplications []MatchingApplication `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"matching_applications"`
	MatchingLikes        []MatchingLike        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"matching_likes"`
	Carpools             []Carpool             `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"carpools"`
	CarpoolApplications  []CarpoolApplication  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"carpool_applications"`
	CarpoolLikes         []CarpoolLike         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"carpool_likes"`
}

// Interest 구조체는 유저의 관심사를 나타냅니다.
type Interest struct {
	InterestID int64  `gorm:"primaryKey;autoIncrement" json:"interest_id"`
	Name       string `gorm:"size:100;not null" json:"name"`
	Users      []User `gorm:"many2many:users_interests" json:"users"`
}
