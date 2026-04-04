package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
    Admin          Role = "admin"
    Doctor         Role = "doctor"
    Surgeon        Role = "surgeon"
    Nurse          Role = "nurse"
    HeadNurse      Role = "head_nurse"
    Receptionist   Role = "receptionist"
    LabTechnician  Role = "lab_technician"
    Pharmacist     Role = "pharmacist"
    ERDoctor       Role = "er_doctor"
    ERNurse        Role = "er_nurse"
    ITSupport      Role = "it_support"
    RobotOperator  Role = "robot_operator"
)
type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	FullName   string    `gorm:"size:255;not null" json:"full_name"`
	Email      string    `gorm:"size:255;unique;not null" json:"email"`
	Phone      string    `gorm:"size:20;uniqueIndex;not null" json:"phone"`
	PasswordHash string  `gorm:"type:text;not null" json:"-"`
	Role       Role      `gorm:"type:varchar(50);not null;index" json:"role"`
	HospitalID uuid.UUID `gorm:"type:uuid;not null;index" json:"hospital_id"`
	DepartmentID *uuid.UUID `gorm:"type:uuid;index" json:"department_id,omitempty"`
	Specialty    string     `gorm:"size:100" json:"specialty,omitempty"` 
	IsActive           bool          `gorm:"default:true" json:"is_active"`
	LastPasswordChange *time.Time    `json:"last_password_change,omitempty"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
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
type UpdateUserInput struct {
	FullName     *string     `json:"full_name"`
	Phone        *string     `json:"phone"`
	Role         *Role `json:"role"`
	DepartmentID *uuid.UUID  `json:"department_id"`
}