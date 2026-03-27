package model

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)
type Hospital struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	// Deployment Info
    TINNumber     string    // For our company records
    ContactPerson string    // Who we deal with at the hospital
    ContactPhone  string
    
    // Robotics/SLAM Config
    // This tells our system which Gazebo/Physical map to load
    MapID         string    
    DockingStationX float64 // Home position for the robot
    DockingStationY float64
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewHospital(name, address, tinnumber, contactperson, contactphone string) (*Hospital, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if address == "" {
		return nil, errors.New("address is required")
	}
    slug := strings.ToLower(strings.ReplaceAll(name, " ", "-"))
	return &Hospital{
		ID:        uuid.New(),
		Name:      name,
		Slug:      slug,
		TINNumber: tinnumber,
		ContactPerson: contactperson,
		ContactPhone: contactphone,
		MapID: "default_map", // Default map, can be updated later
		DockingStationX: 0.0, // Default docking station coordinates
		DockingStationY: 0.0,
		Address:   address,
		CreatedAt: time.Now(),
	}, nil
}