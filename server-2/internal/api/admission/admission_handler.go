package admission

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/middleware"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/admission"
)

func AdmissionHandler(w http.ResponseWriter, r *http.Request) {
	hospitalID, ok := r.Context().Value(middleware.HospitalIDKey).(uuid.UUID)
	role, Rok := r.Context().Value(middleware.RoleKey).(string)

	if !ok || !Rok {
		http.Error(w, "missing hospital id or role", http.StatusBadRequest)
		return
	}

	allowedRoles := map[string]bool{
		"receptionist": true,
		"admin":        true,
	}

	if !allowedRoles[role] {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	var input model.CreateAdmissionRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	admission, err := admission.AdmissionService(input, hospitalID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(admission)
}

func GetAssignedPatientsHandler(w http.ResponseWriter, r *http.Request) {
	staffID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	patients, err := admission.GetAssignedPatientsService(staffID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patients)
}

func GetAdmittedPatientVitalsHandler(w http.ResponseWriter, r *http.Request) {
	hospitalID, ok := r.Context().Value(middleware.HospitalIDKey).(uuid.UUID)
	userID, ok2 := r.Context().Value(middleware.UserIDKey).(uuid.UUID)

	if !ok || !ok2 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vitals, err := admission.GetAdmittedPatientVitalsService(
		hospitalID,
		userID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vitals)
}