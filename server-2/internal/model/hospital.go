package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)
type Hospital struct {
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name string    `gorm:"type:text;not null" json:"name"`
	Slug string    `gorm:"uniqueIndex;not null" json:"slug"`
	Secret        string    `gorm:"size:64;not null" json:"-"`
	Type string `gorm:"size:50" json:"type"` 
	TINNumber     string `gorm:"uniqueIndex;size:100" json:"tin_number"`
	ContactPerson string `gorm:"size:255" json:"contact_person"`
	ContactPhone  string `gorm:"size:20" json:"contact_phone"`
	Address    string  `gorm:"type:text" json:"address"`
	Country    string  `gorm:"size:100" json:"country"`
	City       string  `gorm:"size:100" json:"city"`
	State      string  `gorm:"size:100" json:"state"`
	PostalCode string  `gorm:"size:20" json:"postal_code"`
	Latitude   float64 `gorm:"type:decimal(10,6)" json:"latitude"`
	Longitude  float64 `gorm:"type:decimal(10,6)" json:"longitude"`
	OwnerID  uuid.UUID `gorm:"type:uuid" json:"owner_id"`
	IsActive bool      `gorm:"default:true" json:"is_active"`
	Plan      string `gorm:"size:50" json:"plan"` 
	MaxUsers  int64  `gorm:"type:bigint;default:10" json:"max_users"`
	MaxRobots int64  `gorm:"type:bigint;default:1" json:"max_robots"`
	TotalUsers int64 `gorm:"type:bigint;default:0" json:"total_users"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
type HospitalRequest struct {
	Name          string `json:"name"`
	TINNumber     string `json:"tin_number"`
	ContactPerson string `json:"contact_person"`
	ContactPhone  string `json:"contact_phone"`
	Address       string `json:"address"`
	AdminFullName string `json:"admin_full_name"`
    AdminEmail    string `json:"admin_email"`
    AdminPassword string `json:"admin_password"`
    AdminPhone    string `json:"admin_phone"`
}