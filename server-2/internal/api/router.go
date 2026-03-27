package api

import (
	"github.com/gorilla/mux"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/api/auth"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/middleware"
)
func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware)
    r.HandleFunc("/signup", auth.SignupHandler).Methods("POST")
	r.HandleFunc("/login", auth.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", auth.LogoutHandler).Methods("POST")
	return  r
}