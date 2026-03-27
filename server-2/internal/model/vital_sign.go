package model

import (
	"time"

	"github.com/google/uuid"
)

type VitalSign struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patient_id"`
	RobotID   uuid.UUID `gorm:"type:uuid;not null;index" json:"robot_id"`

	// Vital Measurements
	HeartRate   float64 `gorm:"not null" json:"heart_rate"`   // bpm
	SpO2        float64 `gorm:"not null" json:"spo2"`         // %
	Temperature float64 `gorm:"not null" json:"temperature"`  // °C

	// Timestamp from robot (important for latency handling)
	MeasuredAt time.Time `gorm:"index;not null" json:"measured_at"`

	// Server timestamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}