package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// type Hospital struct {
// 	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
// 	Name string    `gorm:"type:text;not null" json:"name"`
// 	Slug string    `gorm:"uniqueIndex;not null" json:"slug"`
// 	TINNumber     string `gorm:"uniqueIndex;size:100" json:"tin_number"`
// 	ContactPerson string `gorm:"size:255" json:"contact_person"`
// 	ContactPhone  string `gorm:"size:20" json:"contact_phone"`
// 	TotalUsers int64 `gorm:"type:bigint;default:1" json:"total_users"`
// 	MapID           string  `gorm:"size:100;default:'default_map'" json:"map_id"`
// 	DockingStationX float64 `gorm:"type:decimal(10,6);default:0" json:"docking_station_x"`
// 	DockingStationY float64 `gorm:"type:decimal(10,6);default:0" json:"docking_station_y"`

// 	Address   string    `gorm:"type:text" json:"address"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// 	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
// }
type Hospital struct {
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name string    `gorm:"type:text;not null" json:"name"`
	Slug string    `gorm:"uniqueIndex;not null" json:"slug"`

	// Classification
	Type string `gorm:"size:50" json:"type"` // general, specialized, teaching, clinic

	// Legal & Contact Info
	TINNumber     string `gorm:"uniqueIndex;size:100" json:"tin_number"`
	ContactPerson string `gorm:"size:255" json:"contact_person"`
	ContactPhone  string `gorm:"size:20" json:"contact_phone"`

	// Address & Location (important for real systems)
	Address    string  `gorm:"type:text" json:"address"`
	Country    string  `gorm:"size:100" json:"country"`
	City       string  `gorm:"size:100" json:"city"`
	State      string  `gorm:"size:100" json:"state"`
	PostalCode string  `gorm:"size:20" json:"postal_code"`
	Latitude   float64 `gorm:"type:decimal(10,6)" json:"latitude"`
	Longitude  float64 `gorm:"type:decimal(10,6)" json:"longitude"`

	// Ownership & Control
	OwnerID  uuid.UUID `gorm:"type:uuid" json:"owner_id"`
	IsActive bool      `gorm:"default:true" json:"is_active"`

	// Subscription / Limits (for advanced system / SaaS readiness)
	Plan      string `gorm:"size:50" json:"plan"` // basic, premium, enterprise
	MaxUsers  int64  `gorm:"type:bigint;default:10" json:"max_users"`
	MaxRobots int64  `gorm:"type:bigint;default:1" json:"max_robots"`

	// Stats (optional but useful)
	TotalUsers int64 `gorm:"type:bigint;default:0" json:"total_users"`

	// Timestamps
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