package invitation

import (
	"time"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	invitationRepo "github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/invitation"
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
    return invitation, err
}

// func  AcceptAndRegisterStaff(token, name, password, phone string) (model.User, error) {
//     var newUser model.User

//     // Start Transaction
//     err := invitation.Transaction(func(txRepo *invitation.InvitationRepository) error {
//         // 1. Find the invite
//         invite, err := txRepo.GetInviteByToken(token)
//         if err != nil || invite.IsAccepted || time.Now().After(invite.ExpiresAt) {
//             return fmt.Errorf("invitation is no longer valid")
//         }

//         // 2. Create the User
//         hashedPassword, _ := utils.HashPassword(password)
//         newUser = model.User{
//             ID:           uuid.New(),
//             HospitalID:   invite.HospitalID, // Scoped to the same hospital!
//             Email:        invite.Email,
//             FullName:     name,
//             PasswordHash: hashedPassword,
//             Role:         invite.Role,
//             Phone:        phone,
//         }

//         if err := txRepo.CreateUser(newUser); err != nil {
//             return err
//         }

//         // 3. "Burn" the token
//         return txRepo.MarkInviteAsAccepted(invite.ID)
//     })

//     return newUser, err
// }