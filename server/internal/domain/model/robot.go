package model

import (
	"time"

	"github.com/google/uuid"
)
type Robot struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"` // Serial Number
	HospitalID    uint     `gorm:"not null" json:"hospital_id"`
	CurrentRoomID *uint    `json:"current_room_id"`
	Status        string   `gorm:"default:'active'" json:"status"` // "charging", "moving"
	BatteryLevel  int      `json:"battery_level"`
	LastSeen      time.Time `json:"last_seen"`
}