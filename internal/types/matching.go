package types

import (
	"time"
)

type MatchingService interface {
	GetTopLikeMatchingPosts(limit int) ([]MatchingTopLikesResponseDTO, error)
	CreateMatchingPost(request CreateMatchingPostRequestDTO) error
	GetUserMatchingPosts(userID int64) ([]MatchingPostResponseDTO, error)
	CreateMatchingApplication(request MatchingApplicationRequestDTO, matchingID int64) (*MatchingApplicationResponseDTO, error)
	GetPostsForAI(page int, pageSize int) ([]*MatchingDetailForAI, error)
	GetExampleMatchingPosts() ([]*MatchingDetailForAI, error)
}

// Matching 구조체는 매칭 정보를 나타냅니다.
type Matching struct {
	MatchingID   int64                 `gorm:"primaryKey;autoIncrement"`
	Title        string                `gorm:"size:100;not null"`
	ImageURL     string                `gorm:"size:255"`
	Details      string                `gorm:"type:text"`
	UserNickname string                `gorm:"size:50;not null"`
	UserID       int64                 `gorm:"not null"`
	Destination  string                `gorm:"size:100;not null"`
	Date         string                `gorm:"size:100;not null"`
	StartTime    string                `gorm:"size:100" json:"start_time"`
	EndTime      string                `gorm:"size:100" json:"end_time"`
	Status       string                `gorm:"type:varchar(20);default:'active'"`
	CreatedAt    time.Time             `gorm:"autoCreateTime"`
	UpdatedAt    time.Time             `gorm:"autoUpdateTime"`
	User         User                  `gorm:"foreignKey:UserID;references:UserID"`
	Categories   []Category            `gorm:"foreignKey:MatchingID;constraint:OnDelete:CASCADE;"`
	LikesModel   []MatchingLike        `gorm:"foreignKey:MatchingID;constraint:OnDelete:CASCADE;"`
	Applications []MatchingApplication `gorm:"foreignKey:MatchingID;constraint:OnDelete:CASCADE;"`
	Likes        int                   `gorm:"default:0"`
}

type MatchingApplication struct {
	ApplicationID int64     `gorm:"primaryKey;autoIncrement"`
	UserID        int64     `gorm:"not null" json:"user_id"`
	MatchingID    int64     `gorm:"not null"`
	Description   string    `gorm:"type:text" json:"description"` // 자기소개
	AppliedAt     time.Time `gorm:"autoCreateTime"`
	User          *User     `gorm:"foreignKey:UserID"`            // 포인터로 변경
	Matching      *Matching `gorm:"constraint:OnDelete:CASCADE;"` // 포인터로 변경
}

type Category struct {
	CategoryID int64    `gorm:"primaryKey;autoIncrement"`
	Name       string   `gorm:"size:50;not null"`
	MatchingID int64    `gorm:"not null"`
	Matching   Matching `gorm:"foreignKey:MatchingID;constraint:OnDelete:CASCADE;"`
}

type MatchingLike struct {
	LikeID     int64     `gorm:"primaryKey;autoIncrement"`
	UserID     int64     `gorm:"not null"`
	MatchingID int64     `gorm:"not null"`
	LikedAt    time.Time `gorm:"autoCreateTime"`
	User       User      `gorm:"foreignKey:UserID"`
	Matching   Matching  `gorm:"constraint:OnDelete:CASCADE;"`
}

type MatchingTopLikesResponseDTO struct {
	MatchingID      int    `json:"matching_id"`
	Title           string `json:"title"`
	ImageURL        string `json:"image_url"`
	MatchingDetails string `json:"matching_details"`
	LikesCount      int    `json:"likes_count"`
}

type CreateMatchingPostRequestDTO struct {
	Title        string `json:"title" binding:"required"`
	ImageURL     string `json:"image_url"`
	Destination  string `json:"destination" binding:"required"`
	Details      string `json:"details" binding:"required"`
	UserNickname string `json:"user_nickname" binding:"required"`
	UserID       int64  `json:"user_id" binding:"required"`
	Date         string `json:"date" binding:"required"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
}

type CreateMatchingApplicationRequestDTO struct {
	UserID      int64  `json:"user_id" binding:"required"`     // 지원자 ID
	MatchingID  int64  `json:"matching_id" binding:"required"` // 매칭 ID
	Description string `json:"description" binding:"required"` // 자기소개
}

type MatchingPostResponseDTO struct {
	MatchingID  int64  `json:"matching_id"`
	Title       string `json:"title"`
	Destination string `json:"destination"`
	Date        string `json:"date"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	ImageURL    string `json:"image_url"`
}

type GetUserMatchingPostRequestDTO struct {
	UserID int64 `json:"user_id"`
}

type MatchingApplicationRequestDTO struct {
	UserID      int64  `json:"user_id"`
	Description string `json:"description"`
}

type MatchingApplicationResponseDTO struct {
	ApplicationID int64     `json:"application_id"`
	UserID        int64     `json:"user_id"`
	MatchingID    int64     `json:"matching_id"`
	Description   string    `json:"description"` // 자기소개
	AppliedAt     time.Time `json:"applied_at"`
	User          User      `json:"user"`
}

type MatchingDetailForAI struct {
	MatchingID string   `json:"matching_id"`
	Title      string   `json:"title"`
	Details    string   `json:"details"`
	Categories []string `json:"categories"`
}
