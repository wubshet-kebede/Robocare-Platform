package model

import (
	"time"

	"github.com/google/uuid"
)

type RoomStatus string

const (
	RoomAvailable RoomStatus = "available"
	RoomOccupied  RoomStatus = "occupied"
	RoomCleaning  RoomStatus = "cleaning"
)

type Room struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	HospitalID uuid.UUID `gorm:"type:uuid;not null;index" json:"hospital_id"`
	RoomNumber  string `gorm:"size:50;not null" json:"room_number"`   
	LocationName string `gorm:"size:100" json:"location_name"`      
	X   float64 `json:"x"`
	Y   float64 `json:"y"`
	Yaw float64 `json:"yaw"` 
	Floor int `gorm:"index" json:"floor"`
	Status RoomStatus `gorm:"type:varchar(20);default:'available';index" json:"status"`

	// Optional: Relationships
	// Patients []Patient `gorm:"foreignKey:CurrentRoomID" json:"patients,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}