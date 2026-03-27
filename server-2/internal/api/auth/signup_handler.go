package auth

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/auth"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var input model.HospitalRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
   org, user, err := auth.Signup(input)
   fmt.Fprintf(w, "the input is %v", input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
    response := map[string]interface{}{
		"hospital": org,
		"admin":     user,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}