package admission

import (
	"fmt"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
)

func AdmissionRepository(admission model.Admission) (*model.Admission, error) {
    var existing model.Admission
	err := db.DB.Where("patient_id = ? AND is_active = true", admission.PatientID).First(&existing).Error
	if err == nil {
		return nil, fmt.Errorf("patient with ID '%s' already has an active admission", admission.PatientID)
	}
	if err := db.DB.Create(&admission).Error; err != nil {
		return nil, err
	}
	return &admission, nil
}
