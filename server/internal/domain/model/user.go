package model

import (
	"github.com/google/uuid"
)
type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	FullName   string   `gorm:"size:255;not null" json:"full_name"`
	Email      string   `gorm:"uniqueIndex;not null" json:"email"`
	Password   string   `json:"-"` // Hashed password, never exported to JSON
	Role       string   `gorm:"size:50;not null" json:"role"` // "doctor", "staff", "admin"
	HospitalID uuid.UUID `gorm:"type:uuid;not null" json:"hospital_id"`
	// Hospital   Hospital `gorm:"foreignKey:HospitalID" json:"hospital"`
}
