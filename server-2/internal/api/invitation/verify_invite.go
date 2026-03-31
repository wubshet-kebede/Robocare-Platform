package invitation

import (
	"encoding/json"
	"net/http"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/invitation"
)

func VerifyInvitationHandler(w http.ResponseWriter, r *http.Request) {
    // Get token from URL query: /auth/verify-invite?token=xxxxx
    token := r.URL.Query().Get("token")
    if token == "" {
        http.Error(w, "Token is required", http.StatusBadRequest)
        return
    }

    //  Call service to check if this token is valid/active
    invitation, err := invitation.GetValidInvitation(token)
    if err != nil {
        http.Error(w, "Invalid or expired invitation", http.StatusNotFound)
        return
    }

    // Return the invitation details (Email and Role) 
    // so the frontend can pre-fill the "Email" field for the doctor
    json.NewEncoder(w).Encode(invitation)
}
