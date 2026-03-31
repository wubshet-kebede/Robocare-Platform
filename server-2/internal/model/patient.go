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

// type Patient struct {
// 	ID                   uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
// 	HospitalID           uuid.UUID     `gorm:"type:uuid;not null" json:"hospital_id"`
// 	FullName             string        `gorm:"size:255;not null" json:"full_name"`
// 	CurrentRoomID        uuid.UUID     `gorm:"type:uuid" json:"current_room_id"`
// 	MedicalRecordNumber  string        `gorm:"size:100;unique;not null" json:"medical_record_number"`
// 	BedNumber            string        `gorm:"size:20" json:"bed_number"`
// 	MedicalStatus        MedicalStatus `gorm:"type:varchar(20);default:'unknown'" json:"medical_status"`
// 	DateOfBirth          time.Time     `gorm:"not null" json:"date_of_birth"`
// 	Email 			  string        `gorm:"size:255;unique;not null" json:"email"`
// 	Phone 			  string        `gorm:"size:20;uniqueIndex;not null" json:"phone"`
// 	Address			  string        `gorm:"size:255" json:"address"`
// 	EmergencyContactName  string        `gorm:"size:255" json:"emergency_contact_name"`
// 	EmergencyContactPhone string        `gorm:"size:20" json:"emergency_contact_phone"`
// 	BloodType			  string        `gorm:"size:3" json:"blood_type"`
// 	Allergies			  string        `gorm:"size:255" json:"allergies"`
// 	ChronicConditions	  string        `gorm:"size:255" json:"chronic_conditions"`
// 	Age                  int 		 `gorm:"not null" json:"age"`
// 	Gender			   string        `gorm:"size:20" json:"gender"`
// 	// Room Room `gorm:"foreignKey:CurrentRoomID" json:"room"`
// 	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
// 	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
// }

type Patient struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	HospitalID  uuid.UUID `gorm:"type:uuid;not null;index" json:"hospital_id"`
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