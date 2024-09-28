package routes

import (
	"github.com/gorilla/mux"
	"github.com/sabillahsakti/coindropedia/controllers/airdropcontroller"
	"github.com/sabillahsakti/coindropedia/controllers/authcontroller"
	"github.com/sabillahsakti/coindropedia/middlewares"
)

func SetupRoutes(r *mux.Router) {
	// Auth routes
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	//Airdrop routes
	r.HandleFunc("/airdrop", airdropcontroller.Show).Methods("GET")
	r.HandleFunc("/airdrop/{id}", airdropcontroller.Show).Methods("GET")

	// Protected routes (under /api)
	api := r.PathPrefix("/api").Subrouter()

	// API Airdrop
	api.HandleFunc("/airdrop", airdropcontroller.Create).Methods("POST")
	api.HandleFunc("/airdrop/{id}", airdropcontroller.Delete).Methods("DELETE")

	// Add middleware to the API subrouter for JWT authentication
	api.Use(middlewares.JWTMiddleware)
}
