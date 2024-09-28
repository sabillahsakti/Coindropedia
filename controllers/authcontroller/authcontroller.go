package authcontroller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sabillahsakti/coindropedia/config"
	"github.com/sabillahsakti/coindropedia/helper"
	"github.com/sabillahsakti/coindropedia/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan json dari front end
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}

	defer r.Body.Close()

	//ambil data user berdasarkan uname di database
	var user models.User
	if err := config.DB.Where("username =?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Username atau password salah"}
			helper.ResponseJson(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJson(w, http.StatusInternalServerError, response)
			return
		}
	}

	//Pengecekan password valid atau tidak
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Username atau password salah"}
		helper.ResponseJson(w, http.StatusUnauthorized, response)
		return
	}

	// Proses pembuatan token JWT
	expTime := time.Now().Add(time.Minute * 15)
	claims := &config.JWTClaim{
		Username: user.Username,
		ID:       strconv.Itoa(int(user.ID)),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "evoting",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	//Mendeklarasikan algoritma yang akan digunakan untuk signin
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	// set token ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	response := map[string]string{
		"message": "Login Berhasil",
		"token":   token, // tambahkan token di response JSON
	}
	helper.ResponseJson(w, http.StatusOK, response)
	return
}

func Register(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan json dari front end
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}

	defer r.Body.Close()

	//hash pass menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert ke database
	if err := config.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJson(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Hapus token yg ada di cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "Logout Berhasil"}
	helper.ResponseJson(w, http.StatusOK, response)
	return

}

func CheckAuth(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			helper.ResponseJson(w, http.StatusUnauthorized, map[string]string{"message": "Not authenticated"})
			return
		}
		helper.ResponseJson(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	tokenStr := cookie.Value
	claims := &config.JWTClaim{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})

	if err != nil || !token.Valid {
		helper.ResponseJson(w, http.StatusUnauthorized, map[string]string{"message": "Not authenticated"})
		return
	}

	helper.ResponseJson(w, http.StatusOK, map[string]string{"message": "Authenticated"})
}
