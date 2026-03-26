package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)
type Robot struct {
	ID            uuid.UUID `json:"id"` // Serial Number
	HospitalID    uuid.UUID `json:"hospital_id"`
	CurrentRoomID *uint    `json:"current_room_id"`
	Status        string   `json:"status"` // "charging", "moving"
	BatteryLevel  int      `json:"battery_level"`
	LastSeen      time.Time `json:"last_seen"`
}

func NewRobot(HospitalID uuid.UUID, CurrentRoomID uuid.UUID, BatteryLevel int) (*Robot, error) {
	if HospitalID == uuid.Nil {
		return nil, errors.New("hospital ID is required")
	}
	if CurrentRoomID == uuid.Nil {
		return nil, errors.New("current room ID is required")
	}
	return &Robot{
		ID:            uuid.New(),
		HospitalID:    HospitalID,
		CurrentRoomID: nil,
		Status:        "charging",
		BatteryLevel:  100,
		LastSeen:      time.Now(),
	}, nil
}
