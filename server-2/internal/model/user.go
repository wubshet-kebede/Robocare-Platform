package model

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleDoctor UserRole = "doctor"
	RoleNurse  UserRole = "nurse"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	FullName   string    `gorm:"size:255;not null" json:"full_name"`
	Email      string    `gorm:"size:255;unique;not null" json:"email"`
	PasswordHash   string    `gorm:"type:text;not null" json:"-"`
	Role       UserRole  `gorm:"type:varchar(20);not null;index" json:"role"`
	HospitalID uuid.UUID `gorm:"type:uuid;not null;index" json:"hospital_id"`
	Phone          string    `gorm:"type:varchar(20);uniqueIndex;not null" json:"phone"`

	// Optional relationship
	// Hospital Hospital `gorm:"foreignKey:HospitalID" json:"hospital,omitempty"`

	// Timestamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}