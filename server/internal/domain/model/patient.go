package model

import (
	"github.com/google/uuid"
)


type Patient struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	HospitalID    uint     `gorm:"not null" json:"hospital_id"`
	FullName      string   `gorm:"size:255;not null" json:"full_name"`
	CurrentRoomID uint     `json:"current_room_id"`
	Room          Room     `gorm:"foreignKey:CurrentRoomID"`
	MedicalStatus string   `json:"medical_status"` // "stable", "critical"
}