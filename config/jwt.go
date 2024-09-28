package config

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

// Load environment variables from the .env file
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// Fetch JWT_KEY from environment variables
var JWT_KEY = []byte(os.Getenv("JWT_KEY"))

// Define a struct for JWT claims
type JWTClaim struct {
	ID       string `json:"id"`       // Pastikan ada tag JSON
	Username string `json:"username"` // Pastikan ada tag JSON
	jwt.RegisteredClaims
}
