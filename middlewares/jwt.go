package middlewares

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sabillahsakti/coindropedia/config"
	"github.com/sabillahsakti/coindropedia/helper"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ambil token dari header Authorization
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			response := map[string]string{"message": "unauthorize"}
			helper.ResponseJson(w, http.StatusUnauthorized, response)
			return
		}

		// Menghapus prefix "Bearer " jika ada
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		} else {
			response := map[string]string{"message": "unauthorize"}
			helper.ResponseJson(w, http.StatusUnauthorized, response)
			return
		}

		claims := &config.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			log.Println("Error parsing token:", err)
			response := map[string]string{"message": "unauthorize"}
			helper.ResponseJson(w, http.StatusUnauthorized, response)
			return
		}

		if !token.Valid {
			response := map[string]string{"message": "unauthorize"}
			helper.ResponseJson(w, http.StatusUnauthorized, response)
			return
		}

		// Asumsikan claims.UserID adalah string, maka kita konversi ke int
		userID, err := strconv.Atoi(claims.ID)
		if err != nil {
			response := map[string]string{"message": "invalid user id"}

			helper.ResponseJson(w, http.StatusBadRequest, response)
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", userID) // `claims.ID` adalah id user dari token
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
