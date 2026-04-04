package model

import (
	"github.com/google/uuid"
)
type Sequence struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	HospitalID uuid.UUID `gorm:"type:uuid;index"`
	Type       string    `gorm:"size:50"` 
	Year       int
	LastNumber int
}