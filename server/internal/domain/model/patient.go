package model

import (
	"errors"

	"github.com/google/uuid"
)


type Patient struct {
	ID            uuid.UUID `json:"id"`
	HospitalID    uint     `json:"hospital_id"`
	FullName      string   `json:"full_name"`
	CurrentRoomID uint     `json:"current_room_id"`
	Room          Room     `json:"CurrentRoomID"`
	MedicalStatus string   `json:"medical_status"` // "stable", "critical"
}
//  Domain-Driven Design (DDD)
// new patient constructor with the validation logic for the required fields 
func NewPatient(HospitalID uint, FullName string, CurrentRoomID uint, MedicalStatus string) (*Patient, error) {
	if FullName == "" {
		return nil, errors.New("full name is required")
	}
	if MedicalStatus == "" {
		return nil, errors.New("medical status is required")
	}

	return &Patient{
		ID:            uuid.New(),
		HospitalID:    HospitalID,
		FullName:      FullName,
		CurrentRoomID: CurrentRoomID,
		MedicalStatus: MedicalStatus,
	}, nil
}