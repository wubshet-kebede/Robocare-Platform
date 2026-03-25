package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/wubshet-kebede/server/internal/infrastructure"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	infrastructure.Connect()
	

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on :%s\n", port)
	
}