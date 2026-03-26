package model

import (
	"time"

	"github.com/google/uuid"
)

type Vital struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	PatientID   uuid.UUID `gorm:"type:uuid;not null" json:"patient_id"`
	RobotID     uuid.UUID `gorm:"type:uuid" json:"robot_id"`
	HeartRate   float64   `json:"heart_rate"` 
	SpO2        float64   `json:"spo2"`        
	Temperature float64   `json:"temperature"` 
	CreatedAt   time.Time `gorm:"index" json:"created_at"`
}