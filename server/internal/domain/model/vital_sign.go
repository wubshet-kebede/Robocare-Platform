package model

import (
	"time"

	"errors"

	"github.com/google/uuid"
)

type VitalSign struct {
	ID          uuid.UUID `json:"id"`
	PatientID   uuid.UUID `json:"patient_id"`
	RobotID     uuid.UUID `json:"robot_id"`
	HeartRate   float64   `json:"heart_rate"` 
	SpO2        float64   `json:"spo2"`        
	Temperature float64   `json:"temperature"` 
	CreatedAt   time.Time `json:"created_at"`
}

func NewVitalSign(PatientID, RobotID uuid.UUID, HeartRate, SpO2, Temperature float64) (*VitalSign, error) {
	if PatientID == uuid.Nil {
		return nil, errors.New("patient ID is required")
	}
	if RobotID == uuid.Nil {
		return nil, errors.New("robot ID is required")
	}
	return &VitalSign{
		ID:          uuid.New(),
		PatientID:   PatientID,
		RobotID:     RobotID,
		HeartRate:   HeartRate,
		SpO2:        SpO2,
		Temperature: Temperature,
		CreatedAt:   time.Now(),
	}, nil
}
