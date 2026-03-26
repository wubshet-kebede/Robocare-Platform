package model

import (
	"errors"

	"github.com/google/uuid"
)

type Room struct {
	ID         uuid.UUID ` json:"id"`
	HospitalID uuid.UUID ` json:"hospital_id"`
	RoomNumber string   `json:"room_number"` // "Bed-01"
	X          float64  `json:"x_coordinate"` // SLAM X
	Y          float64  `json:"y_coordinate"` // SLAM Y
	Yaw        float64  `json:"yaw"`          // Orientation
	Status     string   `json:"status"` 
}

func NewRoom(HospitalID uuid.UUID, RoomNumber string, X, Y, Yaw float64, Status string) (*Room, error) {
	if HospitalID == uuid.Nil {
		return nil, errors.New("hospital ID is required")
	}
	if RoomNumber == "" {
		return nil, errors.New("room number is required")
	}
	if Status == "" {
		return nil, errors.New("status is required")
	}
	return &Room{
		ID:         uuid.New(),
		HospitalID: HospitalID,
		RoomNumber: RoomNumber,
		X:          X,
		Y:          Y,
		Yaw:        Yaw,
		Status:     Status,
	}, nil
}