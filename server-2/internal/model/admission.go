package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)
type UrgencyLevel string
const (
	UrgencyNormal  UrgencyLevel = "Normal"
	UrgencyHigh    UrgencyLevel = "Urgent"
	UrgencyEmergency UrgencyLevel = "Emergency"
)
type AdmissionStatus string 
const (
	StatusAdmitted  AdmissionStatus = "Admitted"
	StatusDischarged AdmissionStatus = "Discharged"
	StatusWaiting AdmissionStatus = "Waiting"
	StatusDeceased AdmissionStatus = "Deceased"

)

type Admission struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patient_id"`
	HospitalID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_hospital_mrn" json:"hospital_id"`

    AssignedDoctorID uuid.UUID `gorm:"type:uuid;not null;index" json:"assigned_doctor_id"`
	RoomID           uuid.UUID `gorm:"type:uuid;index" json:"room_id"`
	BedNumber        string    `gorm:"size:20" json:"bed_number"`

	MedicalRecordNumber string        `gorm:"size:100;uniqueIndex:idx_hospital_mrn;not null" json:"medical_record_number"`
	Status             MedicalStatus `gorm:"type:varchar(20);default:'unknown'" json:"status"`
	ReasonForAdmission string        `gorm:"type:text" json:"reason_for_admission"`
	Diagnosis          string        `gorm:"type:text" json:"diagnosis"`
	Urgency            UrgencyLevel `gorm:"type:varchar(20);default:'Normal'" json:"urgency"`
	AdmissionStatus    AdmissionStatus `gorm:"type:varchar(20);default:'Waiting'" json:"admission_status"`

	AdmissionDate  time.Time `json:"admission_date"`
	DischargeDate  *time.Time `json:"discharge_date,omitempty"`
	IsActive       bool      `gorm:"default:true;index" json:"is_active"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
type CreateAdmissionRequest struct {
	PatientID          uuid.UUID `json:"patient_id"`
	AssignedDoctorID   uuid.UUID `json:"assigned_doctor_id"`
	RoomID             uuid.UUID `json:"room_id"`
	BedNumber          string    `json:"bed_number"`
	Diagnosis          string    `json:"diagnosis"`
	ReasonForAdmission string    `json:"reason_for_admission"`
	Urgency            string    `json:"urgency"`
	AdmissionDate      time.Time `json:"admission_date"`
	AdmissionStatus    string    `json:"admission_status"`
	Status			 string    `json:"status"`
}
type AssignedPatientResponse struct {
	PatientID          uuid.UUID `json:"patient_id"`
	FullName           string    `json:"full_name"`
	Gender             string    `json:"gender"`
	DateOfBirth        time.Time `json:"date_of_birth"`
	RoomID             uuid.UUID `json:"room_id"`

	Diagnosis          string    `json:"diagnosis"`
	AdmissionStatus    string    `json:"admission_status"`
	Urgency            string    `json:"urgency"`

	RoomNumber         string    `json:"room_number"`
	AssignedDoctorName string    `json:"assigned_doctor_name"`
	HeartRate *float64 `json:"heart_rate"`
    SpO2 *float64 `json:"spo2"`
    Temperature *float64 `json:"temperature"`
    // SystolicBP *int `json:"systolic_bp"`
    // DiastolicBP *int `json:"diastolic_bp"`
    MeasuredAt *time.Time `json:"measured_at"`
}