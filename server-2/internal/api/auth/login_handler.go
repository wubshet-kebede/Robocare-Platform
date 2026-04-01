package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/auth"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var input model.UserLoginInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	token, err := auth.Login(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token["access_token"],
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  time.Now().Add(4 * time.Hour),
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    token["refresh_token"],
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  time.Now().Add(4 * time.Hour),
	})
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
	json.NewEncoder(w).Encode(map[string]interface{}{
	"message":       "Login successful",
	"access_token":  token["access_token"],
	"refresh_token": token["refresh_token"],
})
}
