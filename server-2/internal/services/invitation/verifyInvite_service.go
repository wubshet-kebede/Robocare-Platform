package invitation

import (
	"time"

	"errors"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/invitation"
)
var (
    ErrInvitationNotFound = errors.New("invitation not found")
    ErrInvitationAccepted = errors.New("invitation already accepted")
    ErrInvitationExpired  = errors.New("invitation expired")
)
func GetValidInvitation(token string) (model.Invitation, error) {
    invitation, err := invitation.FindInvitationByToken(token)
    if err != nil {
        return model.Invitation{}, ErrInvitationNotFound
    }

    if invitation.IsAccepted {
        return model.Invitation{}, ErrInvitationAccepted
    }

    if time.Now().After(invitation.ExpiresAt) {
        return model.Invitation{}, ErrInvitationExpired
    }

    return invitation, nil
    
}