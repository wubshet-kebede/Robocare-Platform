package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoomStatus string

const (
	RoomAvailable RoomStatus = "available"
	RoomOccupied  RoomStatus = "occupied"
	RoomCleaning  RoomStatus = "cleaning"
)

type Room struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	HospitalID uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_hospital_room" json:"hospital_id"`
	RoomNumber string    `gorm:"size:50;not null;uniqueIndex:idx_hospital_room" json:"room_number"`

	LocationName string `gorm:"size:100" json:"location_name"`
    X            float64 `gorm:"type:double precision"`
    Y            float64 `gorm:"type:double precision"`
    Yaw          float64 `gorm:"type:double precision"` 
    Floor        int     `gorm:"index" json:"floor"`

	Status RoomStatus `gorm:"type:varchar(20);default:'available';index" json:"status"`

	Capacity     int         `gorm:"default:1" json:"capacity"`
	DepartmentID *uuid.UUID  `gorm:"type:uuid;index" json:"department_id,omitempty"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}