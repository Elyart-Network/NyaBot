package models

type PluginConfig struct {
	ID       int64  `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	Name     string `json:"name" gorm:"not null;unique"`
	Type     string `json:"type" gorm:"not null"`
	Config   string `json:"config" gorm:"not null"`
	Addition string `json:"addition" gorm:"default:null"`
}
