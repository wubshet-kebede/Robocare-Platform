package invitation

import (
	"fmt"
	"time"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/invitation"
)
func GetValidInvitation(token string) (model.Invitation, error) {
    // Ask the repo to find the invitation by the token string
    invitation, err := invitation.FindInvitationByToken(token)
    if err != nil {
        return model.Invitation{}, fmt.Errorf("invitation not found")
    }

    //  Is it already used?
    if invitation.IsAccepted {
        return model.Invitation{}, fmt.Errorf("invitation has already been accepted")
    }

    //  Is it expired?
    if time.Now().After(invitation.ExpiresAt) {
        return model.Invitation{}, fmt.Errorf("invitation link has expired")
    }

    // return the invitation data to the Handler
    return invitation, nil
}