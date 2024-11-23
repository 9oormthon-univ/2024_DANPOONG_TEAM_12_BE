// types 패키지 정의
package types

import "time"

// UsersService 인터페이스 정의
type UsersService interface {
	//RegisterOrLogin(kakaoUser *KakaoUserResponse) (*User, error)
	SignUp(request SignUpRequest) (*User, error)
	Login(request LoginRequest) (*User, error)
}

type User struct {
	UserID       int64     `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username     string    `gorm:"size:50;not null;unique" json:"username"`
	Nickname     string    `gorm:"size:50;not null" json:"nickname"`
	Email        string    `gorm:"size:255;unique" json:"email"`
	Password     string    `gorm:"size:255;not null" json:"-"`
	ProfileImage string    `gorm:"size:512" json:"profile_image"`
	UserDetails  string    `gorm:"type:text" json:"user_details"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Interests    []string  `gorm:"-" json:"interests"`
	//Interests            []Interest            `gorm:"many2many:user_interest;constraint:OnDelete:CASCADE;" json:"interests"`
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

type SignUpRequest struct {
	Username     string   `json:"username" binding:"required"`
	Nickname     string   `json:"nickname" binding:"required"`
	Email        string   `json:"email" binding:"required,email"`
	Password     string   `json:"password" binding:"required,min=6"`
	UserDetails  string   `json:"user_details"`
	ProfileImage string   `json:"profile_image"`
	Interests    []string `json:"interests"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
