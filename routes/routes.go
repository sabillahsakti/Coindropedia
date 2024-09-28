package routes

import (
	"github.com/gorilla/mux"
	"github.com/sabillahsakti/coindropedia/controllers/airdropcontroller"
	"github.com/sabillahsakti/coindropedia/controllers/authcontroller"
	"github.com/sabillahsakti/coindropedia/controllers/favoritecontroller"
	"github.com/sabillahsakti/coindropedia/middlewares"
)

func SetupRoutes(r *mux.Router) {
	// Auth routes
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	//Airdrop routes
	r.HandleFunc("/airdrop", airdropcontroller.GetAll).Methods("GET")
	r.HandleFunc("/airdrop/{id}", airdropcontroller.GetByID).Methods("GET")

	//Favorite routes
	r.HandleFunc("/favorite", favoritecontroller.GetAll).Methods("GET")

	// Protected routes (under /api)
	api := r.PathPrefix("/api").Subrouter()

	// API Airdrop
	api.HandleFunc("/airdrop", airdropcontroller.Create).Methods("POST")
	api.HandleFunc("/airdrop/{id}", airdropcontroller.Update).Methods("PUT")
	api.HandleFunc("/airdrop/{id}", airdropcontroller.Delete).Methods("DELETE")

	// API Favorite
	api.HandleFunc("/favorite", favoritecontroller.GetByID).Methods("GET")
	api.HandleFunc("/favorite", favoritecontroller.Create).Methods("POST")
	api.HandleFunc("/favorite/{id}", favoritecontroller.Delete).Methods("DELETE")

	// Add middleware to the API subrouter for JWT authentication
	api.Use(middlewares.JWTMiddleware)
}
