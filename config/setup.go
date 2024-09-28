package config

import (
	"log"

	"github.com/sabillahsakti/coindropedia/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=sabillahsakti password=password dbname=coindropedia port=5432"
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err)
	}

	DB = db

	// Melakukan migrasi otomatis
	err = db.AutoMigrate(&models.Airdrop{}, &models.Favorite{}, &models.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	} else {
		log.Println("Migration berhasil")
	}
}
