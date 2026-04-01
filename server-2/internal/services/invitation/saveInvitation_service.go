package invitation

import (
	"fmt"
	"time"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/features"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/invitation"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/utils"
)

func AcceptAndRegisterStaff(req model.CompleteProfileRequest) (model.User, error) {
    invite, err := invitation.FindInvitationByToken(req.Token)
    if err != nil {
        return model.User{}, ErrInvitationNotFound
    }
    if invite.IsAccepted {
        return model.User{}, ErrInvitationAccepted
    }
    if time.Now().After(invite.ExpiresAt) {
        return model.User{}, ErrInvitationExpired
    }

    hashedPassword, err := utils.HashPassword(req.Password)
    if err != nil {
        return model.User{}, fmt.Errorf("password processing failed: %w", err)
    }

    // 4. Call repo to create the user
    user, err := features.CreateUserFromInvitation(invite, req.FullName, hashedPassword, req.Phone, req.Specialty)
    if err != nil {
        return model.User{}, fmt.Errorf("creating user failed: %w", err)
    }

    
    return user, nil
}
  