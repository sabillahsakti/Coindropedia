package models

import "time"

// Airdrop Model
type Airdrop struct {
	ID          int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string     `gorm:"type:varchar(100);not null" json:"title"`
	Description string     `gorm:"type:varchar(255);not null" json:"description"` // Deskripsi mungkin lebih panjang
	Status      string     `gorm:"type:varchar(20);not null" json:"status"`       // 'ongoing', 'finished'
	ImageURL    string     `gorm:"type:varchar(255);not null" json:"image_url"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"` // Otomatis diisi saat membuat
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"` // Otomatis diisi saat update
	Favorites   []Favorite `gorm:"foreignKey:AirdropID"`             // Relasi dengan Favorit
}
