package models

import "time"

type OrganizationPosition struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	PositionName  string    `gorm:"size:255;not null" json:"position_name"`
	PositionLevel int       `gorm:"not null" json:"position_level"`
	PositionType  string    `gorm:"type:enum('ketua','wakil','sekretaris','bendahara','bidang');not null" json:"position_type"`
	ParentID      *uint     `json:"parent_id,omitempty"`
	OrderNumber   int       `gorm:"default:0" json:"order_number"`
	IsActive      bool      `gorm:"default:true" json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (OrganizationPosition) TableName() string {
	return "organization_positions"
}

type BoardMember struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	PositionID   uint      `gorm:"not null" json:"position_id"`
	Position     *OrganizationPosition `gorm:"foreignKey:PositionID" json:"position,omitempty"`
	Name         string    `gorm:"size:255;not null" json:"name"`
	Title        string    `gorm:"size:255" json:"title,omitempty"`
	Photo        string    `gorm:"size:500" json:"photo,omitempty"`
	Bio          string    `gorm:"type:text" json:"bio,omitempty"`
	Email        string    `gorm:"size:255" json:"email,omitempty"`
	Phone        string    `gorm:"size:20" json:"phone,omitempty"`
	SocialMedia  string    `gorm:"type:json" json:"social_media,omitempty"`
	PeriodStart  int       `gorm:"type:year;not null" json:"period_start"`
	PeriodEnd    int       `gorm:"type:year;not null" json:"period_end"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
	OrderNumber  int       `gorm:"default:0" json:"order_number"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (BoardMember) TableName() string {
	return "board_members"
}

type Department struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	HeadName    string    `gorm:"size:255" json:"head_name,omitempty"`
	OrderNumber int       `gorm:"default:0" json:"order_number"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Department) TableName() string {
	return "departments"
}
