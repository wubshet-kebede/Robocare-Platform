package model

import (
	"errors"
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
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	FullName   string    `gorm:"size:255;not null" json:"full_name"`
	Email      string    `gorm:"size:255;unique;not null" json:"email"`
	Password   string    `gorm:"size:255;not null" json:"-"` // never exposed
	Role       UserRole  `gorm:"type:varchar(20);not null;index" json:"role"`
	HospitalID uuid.UUID `gorm:"type:uuid;not null;index" json:"hospital_id"`

	// Optional relationship
	// Hospital Hospital `gorm:"foreignKey:HospitalID" json:"hospital,omitempty"`

	// Timestamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
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