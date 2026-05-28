package admission

import (
	"time"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/admission"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/mrn"
)

func AdmissionService(
	input model.CreateAdmissionRequest,
	hospitalID uuid.UUID,
) (model.Admission, error) {

	mrnNumber, err := mrn.GenerateMRN(hospitalID)
	if err != nil {
		return model.Admission{}, err
	}

	admissionPatient := model.Admission{
		PatientID:             input.PatientID,
		HospitalID:            hospitalID,
		AssignedDoctorID:      input.AssignedDoctorID,
		RoomID:                input.RoomID,
		BedNumber:             input.BedNumber,
		Diagnosis:             input.Diagnosis,
		ReasonForAdmission:    input.ReasonForAdmission,
		Urgency:               model.UrgencyLevel(input.Urgency),
		Status:                model.MedicalStatus(input.Status),
		MedicalRecordNumber:   mrnNumber,
		IsActive:              true,
		AdmissionStatus:       model.AdmissionStatus("Admitted"),
		AdmissionDate:         time.Now(),
	}

	savedAdmission, err := admission.AdmissionRepository(
	admissionPatient)
	if err != nil {
		return model.Admission{}, err
	}

	return *savedAdmission, nil
}