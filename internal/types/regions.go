package types

type RegionsService interface {
}

type LocalInfo struct {
	LocalInfoID int    `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"size:100;not null"`
	Features    string `gorm:"type:text"`
	Specialty   string `gorm:"type:text"`
	Symbol      string `gorm:"type:text"`
	Attractions string `gorm:"type:text"`
	Activities  string `gorm:"type:text"`
}
