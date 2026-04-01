package invitation

import (
	"fmt"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
)

func SaveInvitation(invite model.Invitation) error {
    tx := db.DB.Begin() 

    var existing model.Invitation
    if err := tx.
        Where("email = ? AND hospital_id = ? AND is_accepted = false", invite.Email, invite.HospitalID).
        First(&existing).Error; err == nil {
        tx.Rollback() 
        return fmt.Errorf("an active invitation already exists for this email")
    }

    if err := tx.Create(&invite).Error; err != nil {
        tx.Rollback() 
        return err
    }

    return tx.Commit().Error 
}

