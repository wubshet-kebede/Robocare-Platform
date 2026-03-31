package invitation

import (
	"fmt"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
)

func  SaveInvitation(invite model.Invitation) error {
    var existing model.Invitation
    result := db.DB.Where("email = ? AND hospital_id = ? AND accepted_at IS NULL", invite.Email, invite.HospitalID).First(&existing)
    
    if result.Error == nil {
        return fmt.Errorf("an active invitation already exists for this email")
    }
    return db.DB.Create(&invite).Error
}

