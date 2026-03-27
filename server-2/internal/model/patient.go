package model

import (
	"time"

	"github.com/google/uuid"
)

type MedicalStatus string

const (
	StatusUnknown  MedicalStatus = "unknown"
	StatusStable   MedicalStatus = "stable"
	StatusWarning  MedicalStatus = "warning"
	StatusCritical MedicalStatus = "critical"
)

type Patient struct {
	ID                   uuid.UUID     `gorm:"type:uuid;primaryKey" json:"id"`
	HospitalID           uuid.UUID     `gorm:"type:uuid;not null" json:"hospital_id"`
	FullName             string        `gorm:"size:255;not null" json:"full_name"`
	CurrentRoomID        uuid.UUID     `gorm:"type:uuid" json:"current_room_id"`
	MedicalRecordNumber  string        `gorm:"size:100;unique;not null" json:"medical_record_number"`
	BedNumber            string        `gorm:"size:20" json:"bed_number"`
	MedicalStatus        MedicalStatus `gorm:"type:varchar(20);default:'unknown'" json:"medical_status"`
	// Room Room `gorm:"foreignKey:CurrentRoomID" json:"room"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}