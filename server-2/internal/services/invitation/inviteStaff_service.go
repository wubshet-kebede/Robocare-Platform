package invitation

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	invitationRepo "github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/invitation"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/organization"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/mailer"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/utils"
)

func  CreateStaffInvitation(hospitalID uuid.UUID, email string, role model.Role) (model.Invitation, error) {
   
    token := utils.GenerateRandomToken(32) 
    invitation := model.Invitation{
        ID:         uuid.New(),
        HospitalID: hospitalID,
        Email:      email,
        Role:       role,
        Token:      token,
        ExpiresAt:  time.Now().Add(48 * time.Hour), 
    }

    err := invitationRepo.SaveInvitation(invitation)
    go func() {
        
        hospital, err:= organization.GetHospitalByHospitalID(hospitalID)

        if err != nil {
            fmt.Printf("Failed to get hospital by ID: %v\n", err)
            return
        }
        
        err = mailer.SendInvitationEmail(
            invitation.Email, 
            invitation.Token, 
            string(invitation.Role), 
            hospital.Name,
        )
        
        if err != nil {
            
            fmt.Printf("Failed to send invite email to %s: %v\n", invitation.Email, err)
        }
    }()
    return invitation, err
}

