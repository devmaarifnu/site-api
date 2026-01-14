package models

import "time"

// EditorialTeam represents editorial team member
type EditorialTeam struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"size:255;not null"`
	Position    string    `json:"position" gorm:"size:255;not null"`
	RoleType    string    `json:"role_type" gorm:"type:enum('pemimpin_redaksi','wakil_pemimpin_redaksi','redaktur_pelaksana','tim_redaksi');not null"`
	Photo       *string   `json:"photo" gorm:"size:500"`
	Bio         *string   `json:"bio" gorm:"type:text"`
	Email       *string   `json:"email" gorm:"size:255"`
	Phone       *string   `json:"phone" gorm:"size:20"`
	OrderNumber int       `json:"order_number" gorm:"default:0"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName overrides the table name
func (EditorialTeam) TableName() string {
	return "editorial_team"
}

// EditorialCouncil represents editorial council member
type EditorialCouncil struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"size:255;not null"`
	Institution string    `json:"institution" gorm:"size:255;not null"`
	Expertise   *string   `json:"expertise" gorm:"size:500"`
	Photo       *string   `json:"photo" gorm:"size:500"`
	Bio         *string   `json:"bio" gorm:"type:text"`
	Email       *string   `json:"email" gorm:"size:255"`
	OrderNumber int       `json:"order_number" gorm:"default:0"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName overrides the table name
func (EditorialCouncil) TableName() string {
	return "editorial_council"
}
