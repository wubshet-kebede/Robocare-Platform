package model

import (
	"errors"

	"github.com/google/uuid"
)

// type User struct {
// 	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
// 	FullName   string   `gorm:"size:255;not null" json:"full_name"`
// 	Email      string   `gorm:"uniqueIndex;not null" json:"email"`
// 	Password   string   `json:"-"` // Hashed password, never exported to JSON
// 	Role       string   `gorm:"size:50;not null" json:"role"` // "doctor", "staff", "admin"
// 	HospitalID uuid.UUID `gorm:"type:uuid;not null" json:"hospital_id"`
// 	// Hospital   Hospital `gorm:"foreignKey:HospitalID" json:"hospital"`
// }
type UserRole string

const (
    RoleAdmin  UserRole = "admin"
    RoleDoctor UserRole = "doctor"
    RoleNurse  UserRole = "nurse"
)
type User struct {
	ID         uuid.UUID     `json:"id"`
	FullName   string   `json:"full_name"`
	Email      string   `json:"email"`
	Password   string   `json:"-"`    // "-" means never send the password to the Frontend (Nuxt)
	Role       UserRole   `json:"role"` // "doctor" or "staff"
	HospitalID uuid.UUID     `json:"hospital_id"`
	// Hospital   Hospital `json:"hospital"` // Nested struct to show hospital details
}

func NewUser(FullName, Email, Password string, Role UserRole, HospitalID uuid.UUID) (*User, error) {
	if FullName == "" {
		return nil, errors.New("full name is required")
	}
	if Email == "" {
		return nil, errors.New("email is required")
	}
	if Password == "" {
		return nil, errors.New("password is required")
	}
	if Role == "" {
		return nil, errors.New("role is required")
	}
	if HospitalID == uuid.Nil {
		return nil, errors.New("hospital ID is required")
	}
	validRole := false
    roles := []UserRole{RoleAdmin, RoleDoctor, RoleNurse}
    for _, r := range roles {
        if Role == r {
            validRole = true
            break
        }
    }
    if !validRole {
        return nil, errors.New("invalid role assigned")
    }
	return &User{
		ID:         uuid.New(),
		FullName:   FullName,
		Email:      Email,
		Password:   Password,
		Role:       Role,
		HospitalID: HospitalID,
	}, nil
}