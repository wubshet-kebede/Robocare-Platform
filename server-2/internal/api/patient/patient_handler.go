package patient

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/middleware"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/patient"
)
func CreatePatientHandler(w http.ResponseWriter, r *http.Request) {
	hospitalID, ok := r.Context().Value(middleware.HospitalIDKey).(uuid.UUID)
	role, Rok := r.Context().Value(middleware.RoleKey).(string)
	userID, Uok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	if !ok || !Rok || !Uok {
		http.Error(w, "missing hospital_id, role, or user_id", http.StatusBadRequest)
		return
	}
	if role != "receptionist" {
		http.Error(w, "forbidden: only receptionists can create patients", http.StatusForbidden)
		return
	}
	var input model.Patient
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	input.HospitalID = hospitalID
	input.CreatedBy = userID
	newPatient, err := patient.CreatePatient(input)
	if err != nil {
		http.Error(w, "Failed to create patient", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(newPatient)
}