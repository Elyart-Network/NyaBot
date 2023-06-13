package models

type Plugin struct {
	ID       int64  `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	Name     string `json:"name" gorm:"not null;unique"`
	Type     string `json:"type" gorm:"not null"`
	Config   string `json:"config" gorm:"not null"`
	Addition string `json:"addition" gorm:"default:null"`
}

type Logging struct {
	ID         int64  `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	Collection string `json:"collection" gorm:"not null"`
	Content    string `json:"content" gorm:"not null"`
}
