package auth

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/utils"
)

func CreateHospitalWithAdmin(input model.HospitalRequest) (model.Hospital, model.User, error) {
	tx := db.DB.Begin()
	if tx.Error != nil {
		return model.Hospital{}, model.User{}, tx.Error
	}
	Slug := utils.GenerateSlug(input.Name)
	hospital := model.Hospital{
		ID: uuid.New(),
		Name:          input.Name,
		Slug:          Slug,
		TINNumber:     input.TINNumber,
		ContactPerson: input.ContactPerson,
		ContactPhone: input.ContactPhone,
		Address: input.Address,
	}
	fmt.Println("Hospital ID:", hospital.ID)
	if err := tx.Create(&hospital).Error; err != nil {
		tx.Rollback()
		return model.Hospital{}, model.User{}, err
	}
	hashedPassword, err := utils.HashPassword(input.AdminPassword)
	if err != nil {
		tx.Rollback()
		return model.Hospital{}, model.User{}, err
	}
	admin := model.User{
		HospitalID: hospital.ID,
		FullName:       input.AdminFullName,
		Email:          input.AdminEmail,
		PasswordHash:   hashedPassword,
		Role:           model.Admin,
		Phone:          input.AdminPhone,
	}
	if err := tx.Create(&admin).Error; err != nil {
		tx.Rollback()
		return model.Hospital{}, model.User{}, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return model.Hospital{}, model.User{}, err
	}
	return hospital, admin, nil
}