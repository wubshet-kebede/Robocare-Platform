package domain

import (
	"github.com/google/uuid"
)

type Room struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	HospitalID uuid.UUID `gorm:"type:uuid;not null" json:"hospital_id"`
	RoomNumber string   `gorm:"size:100;not null" json:"room_number"` // "Bed-01"
	X          float64  `json:"x_coordinate"` // SLAM X
	Y          float64  `json:"y_coordinate"` // SLAM Y
	Yaw        float64  `json:"yaw"`          // Orientation
	Status     string   `gorm:"default:'available'" json:"status"` 
}