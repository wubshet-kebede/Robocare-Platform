package model

import (
	"time"

	"github.com/google/uuid"
)
type VitalSign struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	AdmissionID uuid.UUID `gorm:"type:uuid;not null;index" json:"admission_id"` 
	RobotID     uuid.UUID `gorm:"type:uuid;not null;index" json:"robot_id"`   
	PatientID   uuid.UUID `gorm:"type:uuid;index" json:"patient_id"`  
	HeartRate   float64 `gorm:"not null" json:"heart_rate"`  
	SpO2        float64 `gorm:"not null" json:"spo2"`         
	Temperature float64 `gorm:"not null" json:"temperature"`  
    SystolicBP  int `json:"systolic_bp,omitempty"`
    DiastolicBP int `json:"diastolic_bp,omitempty"` 
	MeasuredAt time.Time `gorm:"index;not null" json:"measured_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}