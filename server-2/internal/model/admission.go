package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Admission struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patient_id"`
	HospitalID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_hospital_mrn" json:"hospital_id"`

	AssignedDoctorID uuid.UUID `gorm:"type:uuid;index" json:"assigned_doctor_id"`
	RoomID           uuid.UUID `gorm:"type:uuid;index" json:"room_id"`
	BedNumber        string    `gorm:"size:20" json:"bed_number"`

	MedicalRecordNumber string        `gorm:"size:100;uniqueIndex:idx_hospital_mrn;not null" json:"medical_record_number"`
	Status             MedicalStatus `gorm:"type:varchar(20);default:'unknown'" json:"status"`
	ReasonForAdmission string        `gorm:"type:text" json:"reason_for_admission"`
	Diagnosis          string        `gorm:"type:text" json:"diagnosis"`

	AdmissionDate  time.Time `json:"admission_date"`
	DischargeDate  *time.Time `json:"discharge_date,omitempty"`
	IsActive       bool      `gorm:"default:true;index" json:"is_active"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}