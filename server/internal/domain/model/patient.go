package model

import (
	"errors"

	"github.com/google/uuid"
)
type MedicalStatus string

const (
    StatusUnknown  MedicalStatus = "unknown"  // Initial state before robot check
    StatusStable   MedicalStatus = "stable"   // Vitals are in normal range
    StatusWarning  MedicalStatus = "warning"  // One vital is slightly off
    StatusCritical MedicalStatus = "critical" // Vitals are dangerous (e.g., SpO2 < 90%)
)

type Patient struct {
	ID            uuid.UUID `json:"id"`
	HospitalID    uuid.UUID     `json:"hospital_id"`
	FullName      string   `json:"full_name"`
	CurrentRoomID uuid.UUID     `json:"current_room_id"`
	MedicalRecordNumber string `json:"medical_record_number"`
	Room          Room     `json:"CurrentRoomID"`
	BedNumber	  string   `json:"bed_number"`
	MedicalStatus MedicalStatus   `json:"medical_status"` // "stable", "critical"
}
//  Domain-Driven Design (DDD)
// new patient constructor with the validation logic for the required fields 
func NewPatient(HospitalID uuid.UUID, FullName string, CurrentRoomID uuid.UUID, MedicalStatus MedicalStatus, MedicalRecordNumber string) (*Patient, error) {
	if FullName == "" {
		return nil, errors.New("full name is required")
	}
	if MedicalStatus == "" {
		return nil, errors.New("medical status is required")
	}
	if MedicalRecordNumber == "" {
		return nil, errors.New("medical record number is required")
	}

	return &Patient{
		ID:            uuid.New(),
		HospitalID:    HospitalID,
		FullName:      FullName,
		CurrentRoomID: CurrentRoomID,
		MedicalStatus: StatusUnknown, // Default to unknown until robot checks vitals
		MedicalRecordNumber: MedicalRecordNumber,
	}, nil
}