package model

import (
	"time"

	"github.com/google/uuid"
)

// type VitalSign struct {
// 	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
// 	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patient_id"`
// 	RobotID   uuid.UUID `gorm:"type:uuid;not null;index" json:"robot_id"`

// 	// Vital Measurements
// 	HeartRate   float64 `gorm:"not null" json:"heart_rate"`   // bpm
// 	SpO2        float64 `gorm:"not null" json:"spo2"`         // %
// 	Temperature float64 `gorm:"not null" json:"temperature"`  // °C

// 	// Timestamp from robot (important for latency handling)
// 	MeasuredAt time.Time `gorm:"index;not null" json:"measured_at"`

// 	// Server timestamps
// 	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
// }


type VitalSign struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	AdmissionID uuid.UUID `gorm:"type:uuid;not null;index" json:"admission_id"` // Link to admission
	RobotID     uuid.UUID `gorm:"type:uuid;not null;index" json:"robot_id"`     // Which robot recorded it

	// Vital Measurements
	HeartRate   float64 `gorm:"not null" json:"heart_rate"`   // bpm
	SpO2        float64 `gorm:"not null" json:"spo2"`         // %
	Temperature float64 `gorm:"not null" json:"temperature"`  // °C
	// Split Blood Pressure for better querying/graphing
    SystolicBP  int `json:"systolic_bp,omitempty"`
    DiastolicBP int `json:"diastolic_bp,omitempty"` // optional

	// Timestamp from robot (important for latency handling)
	MeasuredAt time.Time `gorm:"index;not null" json:"measured_at"`

	// Server timestamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}