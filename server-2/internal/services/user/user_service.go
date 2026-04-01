package user

import (
	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/features"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/utils"
)

func CreateUser(user model.User, password string) (model.User, error) {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return model.User{}, err
	}
	user.PasswordHash = hashed
	createdUser, err := features.CreateUser(&user)

	if err != nil {
		return model.User{}, err
	}
	return createdUser, nil
}
func ListUsers() ([]model.User, error) {
	return features.GetAllUsers()
}
func CheckForUser(userID uuid.UUID) (model.User, error) {
	return features.CheckForUsers(userID)
}
func UpdateUser(input model.UpdateUserInput, hospitalID, userID uuid.UUID) (*model.User, error) {
	existing, err := features.GetUserByID(userID, hospitalID)
	if err != nil {
		return nil, err
	}
	if input.FullName != nil {
		existing.FullName = *input.FullName
	}

	if input.Phone != nil {
		existing.Phone = *input.Phone
	}

	if input.Role != nil {
		existing.Role = *input.Role
	}

	if input.DepartmentID != nil {
		existing.DepartmentID = input.DepartmentID
	}

	err = features.UpdateUser(existing)
	if err != nil {
		return nil, err
	}

	return existing, nil
}






// func DeleteUser(id uuid.UUID, userhosOID uuid.UUID) (model.User, error) {
// 	user, err := features.GetUserByID(id)
// 	if err != nil {
// 		return model.User{}, err
// 	}
// 	if user.HospitalID != userhosOID {
// 		return model.User{}, fmt.Errorf("forbidden: cannot delete user from another hospital")
// 	}
// 	deletedUser, err := features.DeleteUser(id)

// 	if err != nil {
// 		return model.User{}, err
// 	}
// 	return deletedUser, nil
// }

// func GetUser(userID uuid.UUID, userHosID uuid.UUID) (model.User, error) {

// 	user, err := features.GetUserByID(userID)
// 	if err != nil {
// 		return model.User{}, err
// 	}

// 	if user.HospitalID != userHosID {
// 		return model.User{}, fmt.Errorf("forbidden: cannot access user from another hospital")
// 	}

// 	return user, nil
// }

// func GetUsers(hosID uuid.UUID, page, limit int) ([]model.User, int64, error) {
// 	users, total, err := features.GetUsers(hosID, page, limit)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	return users, total, nil
// }
