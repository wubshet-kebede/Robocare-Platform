package auth

import (
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/utils"
)

func CreateHospitalWithAdmin(input model.HospitalRequest) (model.Hospital, model.User, error) {
	var createdOrg model.Hospital
	var createdAdmin model.User
	tx := db.DB.Begin()
	if tx.Error != nil {
		return createdOrg, createdAdmin, tx.Error
	}
	Slug := utils.GenerateSlug(input.Name)
	hospital := model.Hospital{
		Name:          input.Name,
		Slug:          Slug,
		TINNumber:     input.TINNumber,
		ContactPerson: input.ContactPerson,
		ContactPhone: input.ContactPhone,
		Address: input.Address,
	}
	if err := tx.Create(&hospital).Error; err != nil {
		tx.Rollback()
		return createdOrg, createdAdmin, err
	}
	hashedPassword, pwdErr := utils.HashPassword(input.AdminPassword)
	if pwdErr != nil {
		tx.Rollback()
		return createdOrg, createdAdmin, pwdErr
	}
	admin := model.User{
		HospitalID: hospital.ID,
		FullName:       input.AdminFullName,
		Email:          input.AdminEmail,
		PasswordHash:   hashedPassword,
		Role:           model.RoleAdmin,
		Phone:          input.AdminPhone,
	}
	if err := tx.Create(&admin).Error; err != nil {
		tx.Rollback()
		return hospital, createdAdmin, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return hospital, admin, err
	}
	return hospital, admin, nil
}