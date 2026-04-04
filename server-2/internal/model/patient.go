package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicalStatus string

const (
	StatusUnknown  MedicalStatus = "unknown"
	StatusStable   MedicalStatus = "stable"
	StatusWarning  MedicalStatus = "warning"
	StatusCritical MedicalStatus = "critical"
)
type Patient struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	HospitalID  uuid.UUID `gorm:"type:uuid;not null;index" json:"hospital_id"`
	CreatedBy   uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
	FullName    string    `gorm:"size:255;not null" json:"full_name"`
	DateOfBirth time.Time `gorm:"not null" json:"date_of_birth"`
	Gender      string    `gorm:"type:text" json:"gender"`

	Email  string `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Phone  string `gorm:"size:20;uniqueIndex;not null" json:"phone"`
	Address string `gorm:"size:255" json:"address"`

	EmergencyContactName  string `gorm:"type:text" json:"emergency_contact_name"`
	EmergencyContactPhone string `gorm:"type:text" json:"emergency_contact_phone"`

	BloodType         string `gorm:"type:text" json:"blood_type"`
	Allergies         string `gorm:"type:text" json:"allergies"`
	ChronicConditions string `gorm:"type:text" json:"chronic_conditions"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}