package invitation

import (
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
)

func  FindInvitationByToken(token string) (model.Invitation, error) {
    var invitation model.Invitation
    
    // Look for the record where the token matches
    err := db.DB.Where("token = ?", token).First(&invitation).Error
    if err != nil {
        return invitation, err
    }
    
    return invitation, nil
}