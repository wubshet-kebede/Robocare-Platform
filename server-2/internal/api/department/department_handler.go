package department

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/middleware"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/department"
)
func CreateDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	hospitalID, ok := r.Context().Value(middleware.HospitalIDKey).(uuid.UUID)
	userID, Uok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	role, Rok := r.Context().Value(middleware.RoleKey).(string)
	if !ok || !Uok || !Rok {
		http.Error(w, "missing hospital, user id or role", http.StatusBadRequest)
		return
	}
	if role != "admin" {
		http.Error(w, "forbidden: only admins can create departments", http.StatusForbidden)
		return
	}
	
	var input model.Department
	
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
    input.HospitalID = hospitalID
    input.CreatedBy = userID
	createdDept, err := department.CreateDepartment(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdDept)
}
func GetDepartmentsHandler(w http.ResponseWriter, r*http.Request) {
	hospitalID, ok := r.Context().Value(middleware.HospitalIDKey).(uuid.UUID)
    ///role, Rok := r.Context().Value(middleware.RoleKey).(string)
	if !ok  {
		http.Error(w, "missing hospital id ", http.StatusBadRequest)
		return
	}
	depts, err := department.GetDepartments(hospitalID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(depts)
}