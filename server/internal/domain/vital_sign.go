package domain

import (
	"time"
)

type Vital struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	PatientID   uint      `gorm:"index;not null" json:"patient_id"`
	RobotID     string    `gorm:"size:100" json:"robot_id"`
	HeartRate   float64   `json:"heart_rate"` 
	SpO2        float64   `json:"spo2"`        
	Temperature float64   `json:"temperature"` 
	CreatedAt   time.Time `gorm:"index" json:"created_at"`
}