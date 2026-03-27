package model

import (
	"time"

	"errors"

	"github.com/google/uuid"
)
type VitalSign struct {
    ID          uuid.UUID 
    PatientID   uuid.UUID 
    RobotID     uuid.UUID 
    HeartRate   float64   
    SpO2        float64   
    Temperature float64   
    // MeasuredAt is sent from the Robot to handle network delays
    MeasuredAt  time.Time 
}

func NewVitalSign(patientID, robotID uuid.UUID, hr, spo2, temp float64, measuredAt time.Time) (*VitalSign, error) {
    if patientID == uuid.Nil || robotID == uuid.Nil {
        return nil, errors.New("patient and robot IDs are required")
    }

    // Basic hardware sanity check: ensure sensors aren't sending garbage data
    if hr < 0 || spo2 < 0 || temp < 0 {
        return nil, errors.New("vital measurements cannot be negative")
    }

    return &VitalSign{
        ID:          uuid.New(),
        PatientID:   patientID,
        RobotID:     robotID,
        HeartRate:   hr,
        SpO2:        spo2,
        Temperature: temp,
        MeasuredAt:  measuredAt, // Use the actual time of measurement
    }, nil
}
