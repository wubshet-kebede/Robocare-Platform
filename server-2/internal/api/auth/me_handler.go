package auth

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/middleware"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/user"
)

// func MeHandler(w http.ResponseWriter, r *http.Request) {
// 	userID := r.Context().Value(middleware.UserIDKey).(uuid.UUID)

// 	user, err := user.CheckForUser(userID)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode("something went wrong")
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(user)
// }
func  MeHandler(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "unauthorized",
		})
		return
	}

	data, err := user.GetMe(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "failed to fetch user",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}