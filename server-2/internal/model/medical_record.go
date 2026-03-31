package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type MedicalRecord struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	PatientID   uuid.UUID `gorm:"type:uuid;not null;index" json:"patient_id"`       // permanent patient
	AdmissionID uuid.UUID `gorm:"type:uuid;not null;index" json:"admission_id"`    // links to hospital stay

	// Medical Information
	Diagnosis      string `gorm:"type:text" json:"diagnosis,omitempty"`
	TreatmentPlan  string `gorm:"type:text" json:"treatment_plan,omitempty"`
	Medications    string `gorm:"type:text" json:"medications,omitempty"`
	Procedures     string `gorm:"type:text" json:"procedures,omitempty"`
	LabResults     string `gorm:"type:text" json:"lab_results,omitempty"`
	ImagingReports string `gorm:"type:text" json:"imaging_reports,omitempty"`
	Notes          string `gorm:"type:text" json:"notes,omitempty"` // doctor/nurse notes

	// Status Tracking (optional)
	Status MedicalStatus `gorm:"type:varchar(20);default:'unknown'" json:"status"`

	// Audit / Soft Delete
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Timestamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}