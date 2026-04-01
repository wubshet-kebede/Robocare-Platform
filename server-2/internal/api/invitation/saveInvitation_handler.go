package invitation

import (
	"encoding/json"
	"net/http"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/invitation"
)

func AcceptInvitationHandler(w http.ResponseWriter, r *http.Request) {

    var input model.CompleteProfileRequest

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if input.Token == "" || input.Password == "" || len(input.Password) < 8 {
        http.Error(w, "Token and a strong password are required", http.StatusBadRequest)
        return
    }

   
    user, err := invitation.AcceptAndRegisterStaff(input)
    if err != nil {
        
        http.Error(w, err.Error(), http.StatusUnprocessableEntity)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Account created successfully! You can now sign in with your email.",
		"user information": user,
    })
}