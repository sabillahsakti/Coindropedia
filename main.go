package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sabillahsakti/coindropedia/config"
	"github.com/sabillahsakti/coindropedia/routes"
)

func main() {
	config.ConnectDatabase()

	r := mux.NewRouter()

	routes.SetupRoutes(r)

	// Set up CORS options
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Add frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Wrap the router with the CORS middleware
	handler := corsOptions.Handler(r)

	//Start the server
	log.Println("Server Berjalan di port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
