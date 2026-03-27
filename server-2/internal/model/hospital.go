package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Hospital struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name string    `gorm:"size:255;not null" json:"name"`
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