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
	if role != "receptionist" {
		http.Error(w, "forbidden: only receptionist can create admissions", http.StatusForbidden)
		return
	}
	var input model.Admission
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	input.HospitalID = hospitalID
	admission, err := admission.AdmissionService(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(admission)

}