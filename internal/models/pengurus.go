package models

import "time"

// Pengurus represents organization board member
type Pengurus struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Nama           string    `json:"nama" gorm:"size:255;not null"`
	Jabatan        string    `json:"jabatan" gorm:"size:255;not null"`
	Kategori       string    `json:"kategori" gorm:"type:enum('pimpinan_utama','bidang','sekretariat','bendahara');default:'bidang'"`
	Foto           *string   `json:"foto" gorm:"size:500"`
	Bio            *string   `json:"bio" gorm:"type:text"`
	Email          *string   `json:"email" gorm:"size:255"`
	Phone          *string   `json:"phone" gorm:"size:20"`
	PeriodeMulai   int       `json:"periode_mulai" gorm:"type:year;not null"`
	PeriodeSelesai int       `json:"periode_selesai" gorm:"type:year;not null"`
	OrderNumber    int       `json:"order_number" gorm:"default:0"`
	IsActive       bool      `json:"is_active" gorm:"default:true"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// TableName overrides the table name
func (Pengurus) TableName() string {
	return "pengurus"
}

// PengurusResponse represents the formatted response
type PengurusResponse struct {
	ID       uint    `json:"id"`
	Nama     string  `json:"nama"`
	Jabatan  string  `json:"jabatan"`
	Kategori string  `json:"kategori"`
	Foto     *string `json:"foto"`
	Bio      *string `json:"bio"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
	Periode  struct {
		Mulai   int `json:"mulai"`
		Selesai int `json:"selesai"`
	} `json:"periode"`
	Order    int  `json:"order"`
	IsActive bool `json:"is_active"`
}
