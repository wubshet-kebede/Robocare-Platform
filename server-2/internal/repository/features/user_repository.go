package features

import (
	"errors"
	"fmt"

	"github.com/dlclark/regexp2"
	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"gorm.io/gorm"
)

func CreateUser(user *model.User) (model.User, error) {
	
	if err := db.DB.Create(user).Error; err != nil {
		return model.User{}, err
	}
	
	if err := db.DB.Model(&model.Hospital{}).
		Where("id = ?", user.HospitalID).
		UpdateColumn("total_users", gorm.Expr("total_users + ?", 1)).Error; err != nil {
		return model.User{}, err
	}
	return *user, nil
}
func GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := db.DB.Find(&users).Error
	return users, err
}


func GetUserByEmailOrPhone(identifier string) (model.User, error) {
	var user model.User
	emailRegex := regexp2.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`, 0)
	phoneRegex := regexp2.MustCompile(`^\+?[0-9]{7,15}$`, 0)
	isEmail, _ := emailRegex.MatchString(identifier)
	isPhone, _ := phoneRegex.MatchString(identifier)
	var err error
	switch {
	case isEmail:
		
		err = db.DB.Where("email = ?", identifier).First(&user).Error
	case isPhone:
		
		err = db.DB.Where("phone = ?", identifier).First(&user).Error
	default:
		
		return user, errors.New("invalid identifier format: must be a valid email or phone number")
	}

	return user, err
}
func CheckForUsers(userID uuid.UUID) (model.User, error) {
	var user model.User
	if err := db.DB.First(&user, "id = ?", userID).Error; err != nil {

		return model.User{}, err
	}
	return user, nil
}
func UpdateUser(user *model.User, hosID uuid.UUID) (model.User, error) {
	if hosID == uuid.Nil{
		return model.User{}, errors.New("hospital_id is required")
	}
	var existing model.User
	if err := db.DB.Where("id = ? AND hospital_id = ?", user.ID, hosID).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, fmt.Errorf("user not found")
		}
		return model.User{}, err
	}
	if err := db.DB.Model(&existing).Updates(user).Error; err != nil {
		return model.User{}, err
	}
	return existing, nil
}
func DeleteUser(id uuid.UUID) (model.User, error) {
	var user model.User
	if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
		return model.User{}, err
	}

	if err := db.DB.Delete(&user).Error; err != nil {
		return model.User{}, err
	}
	if err := db.DB.Model(&model.Hospital{}).
		Where("id = ?", user.HospitalID).
		UpdateColumn("total_users", gorm.Expr("total_users - ?", 1)).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}
func GetUserByID(id uuid.UUID) (model.User, error) {
	var user model.User
	if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
func GetUsers(hosID uuid.UUID, page, limit int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	offset := (page - 1) * limit

	if err := db.DB.Model(&model.User{}).
		Where("hospital_id = ?", hosID).
		Count(&total).
		Offset(offset).
		Limit(limit).
		Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
