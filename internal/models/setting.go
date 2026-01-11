package models

import "time"

type Setting struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	SettingKey   string    `gorm:"size:100;unique;not null" json:"setting_key"`
	SettingValue string    `gorm:"type:text" json:"setting_value,omitempty"`
	SettingType  string    `gorm:"type:enum('string','text','number','boolean','json');default:'string'" json:"setting_type"`
	SettingGroup string    `gorm:"size:50;default:'general'" json:"setting_group"`
	Description  string    `gorm:"type:text" json:"description,omitempty"`
	IsPublic     bool      `gorm:"default:false" json:"is_public"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (Setting) TableName() string {
	return "settings"
}
