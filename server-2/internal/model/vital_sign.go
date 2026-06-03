package model

import (
	"time"

	"github.com/google/uuid"
)
type VitalSign struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	HospitalID  uuid.UUID `gorm:"type:uuid;not null;index" json:"hospital_id"`
	AdmissionID uuid.UUID `gorm:"type:uuid;not null;index" json:"admission_id"` 
	CommandID *uuid.UUID `gorm:"type:uuid;index"`
	RobotID     uuid.UUID `gorm:"type:uuid;not null;index" json:"robot_id"`   
	PatientID   uuid.UUID `gorm:"type:uuid;index" json:"patient_id"`  
	HeartRate   float64 `gorm:"not null" json:"heart_rate"`  
	SpO2        float64 `gorm:"not null" json:"spo2"`         
	Temperature float64 `gorm:"not null" json:"temperature"` 
	IsRobotEntry bool	`gorm:"default:false" json:"is_robot_entry"`
    SystolicBP  int `json:"systolic_bp,omitempty"`
    DiastolicBP int `json:"diastolic_bp,omitempty"` 
	MeasuredAt time.Time `gorm:"index;not null" json:"measured_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
type AdmittedPatientVitalResponse struct {
	PatientID      uuid.UUID `json:"patient_id"`
	AdmissionID    uuid.UUID `json:"admission_id"`
	PatientName    string    `json:"patient_name"`
	Status string `json:"status"`

	HeartRate      float64   `json:"heart_rate"`
	SpO2           float64   `json:"spo2"`
	Temperature    float64   `json:"temperature"`
	SystolicBP     int       `json:"systolic_bp"`
	DiastolicBP    int       `json:"diastolic_bp"`

	MeasuredAt     time.Time `json:"measured_at"`
}