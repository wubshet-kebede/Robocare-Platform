package api

import (
	"github.com/gorilla/mux"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/api/auth"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/api/department"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/api/invitation"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/api/patient"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/middleware"
)
func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/invite-staff", invitation.InviteStaffHandler).Methods("POST")
	protected.HandleFunc("/departments", department.CreateDepartmentHandler).Methods("POST") 
	protected.HandleFunc("/departments", department.GetDepartmentsHandler).Methods("GET")
	protected.HandleFunc("/patients", patient.CreatePatientHandler).Methods("POST")
    r.HandleFunc("/signup", auth.SignupHandler).Methods("POST")
	r.HandleFunc("/verify-invite", invitation.VerifyInvitationHandler).Methods("GET")
	r.HandleFunc("/accept-invite", invitation.AcceptInvitationHandler).Methods("POST")
	r.HandleFunc("/login", auth.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", auth.LogoutHandler).Methods("POST")
	return  r
}