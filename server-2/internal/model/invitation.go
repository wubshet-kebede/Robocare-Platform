package model

import (
	"time"

	"github.com/google/uuid"
)

type Invitation struct {
    ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    HospitalID  uuid.UUID `gorm:"type:uuid;not null"`
    Email       string    `gorm:"size:255;not null"`
    Role        Role      `gorm:"type:varchar(20);not null"`
    Token       string    `gorm:"uniqueIndex;not null"` // The secret string in the URL
    IsAccepted  bool      `gorm:"default:false"`
    ExpiresAt   time.Time `json:"expires_at"`
    CreatedAt   time.Time
}