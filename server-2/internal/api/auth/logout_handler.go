package auth

import (
	"encoding/json"
	"net/http"
	"time"
)




func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  time.Unix(0, 0), 
		MaxAge:   -1,
	})

	
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}
