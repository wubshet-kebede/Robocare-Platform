package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)



type DepartmentType string

const (
	Clinical      DepartmentType = "clinical"
	Administrative DepartmentType = "administrative"
	Laboratory     DepartmentType = "laboratory"
	Emergency      DepartmentType = "emergency"
	Pharmacy       DepartmentType = "pharmacy"
	Custom          DepartmentType = "custom"
)

type Department struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    HospitalID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_hospital_dept_code" json:"hospital_id"`
	CreatedBy uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
	Name        string         `gorm:"size:100;not null" json:"name"`              
    Code       string    `gorm:"size:10;not null;uniqueIndex:idx_hospital_dept_code" json:"code"`
	Description string         `gorm:"type:text" json:"description,omitempty"`     
	Type     DepartmentType `gorm:"type:department_type_enum;default:'clinical'" json:"type"`
	MaxBeds  int            `gorm:"default:0" json:"max_beds"`                   
	Floor    int            `json:"floor,omitempty"`                             
	HeadID *uuid.UUID `gorm:"type:uuid;index" json:"head_id,omitempty"`    
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}