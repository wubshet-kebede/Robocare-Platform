package main

import (
	"context"
	"log"
	"os"

	"fmt"

	"github.com/joho/godotenv"

	"github.com/wubshet-kebede/server/internal/adapters/outbound/persistence"
	"github.com/wubshet-kebede/server/internal/application/service"
	"github.com/wubshet-kebede/server/internal/domain/model"
	"github.com/wubshet-kebede/server/internal/infrastructure/db"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database :=db.InitDB()
	err := database.AutoMigrate(
		&model.Hospital{},
		// &model.User{},
		// &model.Room{},
		// &model.Patient{},
		// &model.Vital{},
	)
    if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	// 3. HEXAGONAL WIRING (The "Magic" Part)
	// Create the Adapter (Persistence)
	hospitalRepo := persistence.NewHospitalRepository(database)
	
	// Create the Application Service (The Brain)
	// We pass the repository into the service.
	hospitalService := service.NewHospitalService(hospitalRepo)

	ctx := context.Background()
	name := "University of Gondar Specialized Hospital"
	tinNumber := "1234567890"
	address := "Gondar, Ethiopia"
	contactPerson := "Dr. John Doe"
	contactPhone := "0912345678"

	// We call the service, NOT the database directly.
	newHospital, err := hospitalService.RegisterHospital(ctx, name, address, tinNumber, contactPerson, contactPhone)
	if err != nil {
    log.Fatalf("Registration failed: %v", err) // Stop here if there is an error
}
fmt.Printf("Success! Hospital Created: %v\n", newHospital)
	
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on :%s\n", port)
	
}

