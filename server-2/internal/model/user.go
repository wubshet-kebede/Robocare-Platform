package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	Admin  Role = "admin"
	Doctor Role = "doctor"
	Nurse  Role = "nurse"
)

// type User struct {
// 	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
// 	FullName   string    `gorm:"size:255;not null" json:"full_name"`
// 	Email      string    `gorm:"size:255;unique;not null" json:"email"`
// 	PasswordHash   string    `gorm:"type:text;not null" json:"-"`
// 	Role       Role  `gorm:"type:varchar(20);not null;index" json:"role"`
// 	HospitalID uuid.UUID `gorm:"type:uuid;not null;index" json:"hospital_id"`
// 	Phone          string    `gorm:"type:varchar(20);uniqueIndex;not null" json:"phone"`

	
// 	// Timestamps
// 	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
// 	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
// }
type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	FullName   string    `gorm:"size:255;not null" json:"full_name"`
	Email      string    `gorm:"size:255;unique;not null" json:"email"`
	Phone      string    `gorm:"size:20;uniqueIndex;not null" json:"phone"`
	PasswordHash string  `gorm:"type:text;not null" json:"-"`

	Role       Role      `gorm:"type:varchar(50);not null;index" json:"role"`
	HospitalID uuid.UUID `gorm:"type:uuid;not null;index" json:"hospital_id"`

	// Optional: assign to department
	DepartmentID *uuid.UUID `gorm:"type:uuid;index" json:"department_id,omitempty"`
	Specialty    string     `gorm:"size:100" json:"specialty,omitempty"` // for doctors

	// Active & Audit
	IsActive           bool          `gorm:"default:true" json:"is_active"`
	LastPasswordChange *time.Time    `json:"last_password_change,omitempty"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`

	// Timestamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
type UserLoginInput struct {
    Identifier string `json:"identifier"` 
    Password   string `json:"password"`
}

type UserLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type CompleteProfileRequest struct {
    Token     string `json:"token" validate:"required"`
    FullName  string `json:"full_name" validate:"required,min=3"`
    Phone     string `json:"phone" validate:"required"`
    Password  string `json:"password" validate:"required,min=8"`
    Specialty string `json:"specialty,omitempty"` 
}