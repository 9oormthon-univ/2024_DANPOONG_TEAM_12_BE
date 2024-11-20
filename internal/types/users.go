// types 패키지 정의
package types

// UsersService 인터페이스 정의
type UsersService interface {
}

type User struct {
	UserID               int64                 `gorm:"primaryKey;autoIncrement"`                                         // 유저 ID, 자동 증가 및 기본 키
	Username             string                `gorm:"size:50;not null"`                                                 // 유저 이름, 최대 50자, 필수 입력
	Nickname             string                `gorm:"size:50;not null"`                                                 // 닉네임, 최대 50자, 필수 입력
	UserDetails          string                `gorm:"type:text"`                                                        // 유저 상세 정보, 텍스트 타입
	Interests            []Interest            `gorm:"many2many:user_interest;constraint:OnDelete:CASCADE;"`             // 관심사 다대다 관계, 유저 삭제 시 함께 삭제
	Matchings            []Matching            `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE;"` // 작성한 매칭 게시글, 유저 삭제 시 함께 삭제
	MatchingApplications []MatchingApplication `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`                   // 매칭 지원서도 유저가, 유저 삭제 시 함께 삭제
	MatchingLikes        []MatchingLike        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`                   // 매칭 좋아요, 유저 삭제 시 함께 삭제
	Carpools             []Carpool             `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`                   // 작성한 카풀 게시글, 유저 삭제 시 함께 삭제
	CarpoolApplications  []CarpoolApplication  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`                   // 카풀 지원서, 유저 삭제 시 함께 삭제
	CarpoolLikes         []CarpoolLike         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`                   // 카풀 좋아요, 유저 삭제 시 함께 삭제
}

// Interest 구조체는 유저의 관심사를 나타냅니다.
type Interest struct {
	InterestID int64  `gorm:"primaryKey;autoIncrement"` // 관심사 ID, 자동 증가 및 기본 키
	Name       string `gorm:"size:100;not null"`        // 관심사 이름, 최대 100자, 필수 입력
	Users      []User `gorm:"many2many:user_interest"`  // 관심사를 가진 유저들
}
