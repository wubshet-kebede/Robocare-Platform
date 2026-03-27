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
	TINNumber     string `gorm:"uniqueIndex;size:100" json:"tin_number"`
	ContactPerson string `gorm:"size:255" json:"contact_person"`
	ContactPhone  string `gorm:"size:20" json:"contact_phone"`
	MapID           string  `gorm:"size:100;default:'default_map'" json:"map_id"`
	DockingStationX float64 `gorm:"type:decimal(10,6);default:0" json:"docking_station_x"`
	DockingStationY float64 `gorm:"type:decimal(10,6);default:0" json:"docking_station_y"`

	Address   string    `gorm:"type:text" json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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