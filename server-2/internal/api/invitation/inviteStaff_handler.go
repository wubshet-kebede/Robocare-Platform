package invitation

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/middleware"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/invitation"
)

func InviteStaffHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Email string         `json:"email"`
        Role  model.Role `json:"role"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

	 hospitalID , ok := r.Context().Value(middleware.HospitalIDKey).(uuid.UUID)
	 if !ok  {
		http.Error(w, "missing hosppital id", http.StatusBadRequest)
		return
	}

    invite, err := invitation.CreateStaffInvitation(hospitalID, input.Email, input.Role)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Invitation sent successfully",
        "token":   invite.Token, 
    })
}