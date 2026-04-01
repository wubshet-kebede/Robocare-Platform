package user

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/middleware"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/user"
)

func UpdateStaffHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	hospitalID, ok := r.Context().Value(middleware.HospitalIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "missing hospital id", http.StatusBadRequest)
		return
	}

	var input model.UpdateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	updatedUser, err := user.UpdateUser(input, hospitalID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedUser)
}