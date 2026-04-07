package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RobotStatus string

const (
	RobotIdle       RobotStatus = "idle"
	RobotNavigating RobotStatus = "navigating"
	RobotMeasuring  RobotStatus = "measuring"
	RobotCharging   RobotStatus = "charging"
	RobotEmergency  RobotStatus = "emergency_stop"
	RobotOffline    RobotStatus = "offline"
)
type CommandStatus string

const (
	CommandPending   CommandStatus = "pending"
	CommandSent      CommandStatus = "sent"
	CommandRunning   CommandStatus = "running"
	CommandCompleted CommandStatus = "completed"
	CommandFailed    CommandStatus = "failed"
)

type CommandType string

const (
	CommandMeasureVitals CommandType = "measure_vitals"
)


type Robot struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	HospitalID uuid.UUID `gorm:"type:uuid;not null;index" json:"hospital_id"`
	
	SerialNumber string `gorm:"size:100;uniqueIndex;not null" json:"serial_number"`
	CurrentRoomID *uuid.UUID `gorm:"type:uuid;index" json:"current_room_id,omitempty"`
	TargetRoomID  *uuid.UUID `gorm:"type:uuid;index" json:"target_room_id,omitempty"`
	Status        RobotStatus `gorm:"type:varchar(30);default:'idle';index" json:"status"`
	BatteryLevel int    `gorm:"check:battery_level >= 0 AND battery_level <= 100" json:"battery_level"`
	WiFiStrength int    `json:"wifi_strength"` 
	IPAddress    string `gorm:"size:50" json:"ip_address"`
	LastSeen     time.Time `gorm:"index" json:"last_seen"`
	 
	CurrentAdmissionID *uuid.UUID `gorm:"type:uuid;index" json:"current_admission_id,omitempty"`
	MapID string `gorm:"size:100;default:'default_map'" json:"map_id"`
	Floor int    `json:"floor"`
	Zone  string `gorm:"size:50" json:"zone"`
	Model string `gorm:"size:50" json:"model,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CurrentRoom *Room `gorm:"foreignKey:CurrentRoomID" json:"current_room,omitempty"`
	TargetRoom  *Room `gorm:"foreignKey:TargetRoomID" json:"target_room,omitempty"`
}

type RobotCommand struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	RobotID   uuid.UUID `gorm:"type:uuid;not null;index"`

	PatientID   uuid.UUID `gorm:"type:uuid;not null;index"`
	AdmissionID uuid.UUID `gorm:"type:uuid;not null;index"`
	RoomID      uuid.UUID `gorm:"type:uuid;not null;index"`

	Type   CommandType   `gorm:"type:varchar(50);not null"`
	Status CommandStatus `gorm:"type:varchar(30);default:'pending';index"`

	RequestedBy uuid.UUID `gorm:"type:uuid;not null"` 

	RequestedAt time.Time
	StartedAt   *time.Time
	CompletedAt *time.Time

	ErrorMessage string `gorm:"size:255"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
