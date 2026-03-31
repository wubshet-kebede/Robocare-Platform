package invitation

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/invitation"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/utils"
)

func AcceptInvitationHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Token    string `json:"token"`
        FullName string `json:"full_name"`
        Password string `json:"password"`
        Phone    string `json:"phone"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // 1. Call the Service to "Burn" the token and Create the User
    user, err := invitation.AcceptAndRegisterStaff(input.Token, input.FullName, input.Password, input.Phone)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnprocessableEntity)
        return
    }

    // 2. Since they just registered, generate a JWT immediately 
    // so they are "Logged In" right away.
    token, _ := utils.GenerateJWT(user.ID.String(), user.HospitalID.String(), string(user.Role), 24*time.Hour)

    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Welcome to the team!",
        "token":   token,
        "user":    user,
    })
}