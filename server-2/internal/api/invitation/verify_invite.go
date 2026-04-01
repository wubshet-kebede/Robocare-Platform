package invitation

import (
	"encoding/json"
	"net/http"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/invitation"
)

func VerifyInvitationHandler(w http.ResponseWriter, r *http.Request) {
    token := r.URL.Query().Get("token")
    if token == "" {
        http.Error(w, "Token is required", http.StatusBadRequest)
        return
    }

    invite, err := invitation.GetValidInvitation(token)
    if err != nil {
    switch err {
    case invitation.ErrInvitationNotFound:
        http.Error(w, err.Error(), http.StatusNotFound) 
    case invitation.ErrInvitationAccepted:
        http.Error(w, err.Error(), http.StatusConflict) 
    case invitation.ErrInvitationExpired:
        http.Error(w, err.Error(), http.StatusGone) 
    default:
        http.Error(w, "Internal server error", http.StatusInternalServerError)
    }
    return
}

    json.NewEncoder(w).Encode(invite)
}
