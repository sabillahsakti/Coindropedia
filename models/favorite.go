package models

import "time"

// Favorite Model
type Favorite struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int       `gorm:"not null" json:"user_id"`
	AirdropID int       `gorm:"not null" json:"airdrop_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relasi ke tabel lain
	User    User    `gorm:"foreignKey:UserID" json:"user"`
	Airdrop Airdrop `gorm:"foreignKey:AirdropID" json:"airdrop"`
}
