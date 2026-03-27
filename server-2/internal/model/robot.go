package model

import (
	"time"

	"github.com/google/uuid"
)

type RobotStatus string

const (
	RobotIdle       RobotStatus = "idle"
	RobotNavigating RobotStatus = "navigating"
	RobotCharging   RobotStatus = "charging"
	RobotEmergency  RobotStatus = "emergency_stop"
	RobotOffline    RobotStatus = "offline"
)

type Robot struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	HospitalID uuid.UUID `gorm:"type:uuid;not null" json:"hospital_id"`

	// Navigation State
	CurrentRoomID *uuid.UUID `gorm:"type:uuid" json:"current_room_id"`
	TargetRoomID  *uuid.UUID `gorm:"type:uuid" json:"target_room_id"`
	Status        RobotStatus `gorm:"type:varchar(30);default:'idle'" json:"status"`

	// Telemetry (from ROS2)
	BatteryLevel int    `gorm:"check:battery_level >= 0 AND battery_level <= 100" json:"battery_level"`
	WiFiStrength int    `json:"wifi_strength"` // could be RSSI or percentage
	IPAddress    string `gorm:"size:50" json:"ip_address"`

	// Health / Monitoring
	LastSeen time.Time `gorm:"index" json:"last_seen"`

	// Optional relationships (if you have Room model)
	// CurrentRoom *Room `gorm:"foreignKey:CurrentRoomID" json:"current_room,omitempty"`
	// TargetRoom  *Room `gorm:"foreignKey:TargetRoomID" json:"target_room,omitempty"`

	// Timestamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}