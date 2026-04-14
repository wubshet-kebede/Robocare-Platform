package vitals

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/admission"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/robot"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/vitals"
)

func ProcessVitals(hospitalID, robotID, patientID uuid.UUID, vital model.VitalSign) error {
    if hospitalID == uuid.Nil || patientID == uuid.Nil {
        return errors.New("missing hospital or patient ID")
    }    
    patient, err := admission.GetPatientAdmissionID(patientID)
    if err != nil {
        return fmt.Errorf("patient not found")
    }

    if patient.HospitalID != hospitalID {
        return fmt.Errorf("patient does not belong to this hospital")
    }
    robot, err := robot.GetRobotByID(robotID)
    if err != nil {
        return fmt.Errorf("robot not found")
    }

    if robot.HospitalID != hospitalID {
        return fmt.Errorf("robot does not belong to this hospital")
    }
    err = vitals.SaveVitals(hospitalID, patient.ID, robotID, patientID, vital)
    if err != nil {
        log.Println("DB insert failed:", err)
        return err
    }

    log.Printf("Vitals stored → hospital=%s patient=%s robot=%s",
        hospitalID, patientID, robotID)

    return nil
}