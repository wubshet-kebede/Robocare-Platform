package admission

import (
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/admission"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/mrn"
)

func AdmissionService(input model.Admission) (model.Admission, error) {
	if input.MedicalRecordNumber == "" {
		mrn, err := mrn.GenerateMRN(input.HospitalID)
		if err != nil {
			return model.Admission{}, err
		}
		input.MedicalRecordNumber = mrn
	}

	savedAdmission, err := admission.AdmissionRepository(input)
	if err != nil {
		return model.Admission{}, err
	}

	return *savedAdmission, nil
	
}