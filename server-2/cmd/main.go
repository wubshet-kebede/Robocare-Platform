package main

import (
	"log"
	"os"

	// "github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db.Connect()
	// db.Migrate()
	// r := api.SetupRouter()
	// allowedOrigins := []string{"http://localhost:5173"}
	// allowedHeaders := []string{"Content-Type", "Authorization"}
	// allowedMethods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	// corsMiddleware := handlers.CORS(
	// 	handlers.AllowedOrigins(allowedOrigins),
	// 	handlers.AllowedHeaders(allowedHeaders),
	// 	handlers.AllowedMethods(allowedMethods),
	// 	handlers.AllowCredentials(),
	// )

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on :%s\n", port)
	// log.Fatal(http.ListenAndServe(":"+port, corsMiddleware(r)))
}