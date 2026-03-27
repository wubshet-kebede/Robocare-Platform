package model

import (
	"errors"

	"github.com/google/uuid"
)
type RoomStatus string

const (
    RoomAvailable RoomStatus = "available"
    RoomOccupied  RoomStatus = "occupied"
    RoomCleaning  RoomStatus = "cleaning"
)
type Room struct {
    ID         uuid.UUID 
    HospitalID uuid.UUID 
    RoomNumber string    // e.g., "Room 302"
    LocationName string  // e.g., "Bed A" or "Nurse Station"
    
    // SLAM Coordinates from your Gazebo simulation
    X          float64   
    Y          float64   
    Yaw        float64   
    
    // Mapping
    Floor      int       // For multi-floor hospitals
    
    Status     RoomStatus    // "available", "occupied", "restricted"
}

func NewRoom(hospitalID uuid.UUID, roomNumber string, x, y, yaw float64) (*Room, error) {
    if hospitalID == uuid.Nil {
        return nil, errors.New("hospital ID is required")
    }
    if roomNumber == "" {
        return nil, errors.New("room number is required")
    }
    
    return &Room{
        ID:         uuid.New(),
        HospitalID: hospitalID,
        RoomNumber: roomNumber,
        X:          x,
        Y:          y,
        Yaw:        yaw,
        Status:     RoomAvailable,
    }, nil
}