package department

import (
	"fmt"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"gorm.io/gorm"
)
func CreateDepartment(dept model.Department) (*model.Department, error) {
	var existing model.Department

	err := db.DB.
		Where("code = ? AND hospital_id = ?", dept.Code, dept.HospitalID).
		First(&existing).Error

	if err == nil {
		return nil, fmt.Errorf("department code '%s' already exists in this hospital", dept.Code)
	}
	if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = db.DB.Create(&dept).Error
	if err != nil {
		return nil, err
	}

	return &dept, nil
}