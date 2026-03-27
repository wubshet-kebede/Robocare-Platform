package features

import (
	"errors"

	"github.com/dlclark/regexp2"
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
// func GetAllEmployee() ([]model.Employee, error) {
// 	var emps []model.Employee
// 	err := db.DB.Find(&emps).Error
// 	return emps, err
// }


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
// func CheckForEmployee(userID uuid.UUID) (model.Employee, error) {
// 	var user model.Employee
// 	if err := db.DB.First(&user, "id = ?", userID).Error; err != nil {

// 		return model.Employee{}, err
// 	}
// 	return user, nil
// }
// func UpdateEmployee(emp *model.Employee, orgID uuid.UUID) (model.Employee, error) {
// 	if orgID == uuid.Nil{
// 		return model.Employee{}, errors.New("organization_id is required")
// 	}
// 	var existing model.Employee
// 	if err := db.DB.Where("id = ? AND organization_id = ?", emp.ID, orgID).First(&existing).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return model.Employee{}, fmt.Errorf("employee not found")
// 		}
// 		return model.Employee{}, err
// 	}
// 	if err := db.DB.Model(&existing).Updates(emp).Error; err != nil {
// 		return model.Employee{}, err
// 	}
// 	return existing, nil
// }
// func DeleteEmployee(id uuid.UUID) (model.Employee, error) {
// 	var emp model.Employee
// 	if err := db.DB.First(&emp, "id = ?", id).Error; err != nil {
// 		return model.Employee{}, err
// 	}

// 	if err := db.DB.Delete(&emp).Error; err != nil {
// 		return model.Employee{}, err
// 	}
// 	if err := db.DB.Model(&model.Organization{}).
// 		Where("id = ?", emp.OrganizationID).
// 		UpdateColumn("total_employees", gorm.Expr("total_employees - ?", 1)).Error; err != nil {
// 		return model.Employee{}, err
// 	}

// 	return emp, nil
// }
// func GetEmployeeByID(id uuid.UUID) (model.Employee, error) {
// 	var emp model.Employee
// 	if err := db.DB.First(&emp, "id = ?", id).Error; err != nil {
// 		return model.Employee{}, err
// 	}
// 	return emp, nil
// }
// func GetEmployees(orgID uuid.UUID, page, limit int) ([]model.Employee, int64, error) {
// 	var employees []model.Employee
// 	var total int64

// 	offset := (page - 1) * limit

// 	if err := db.DB.Model(&model.Employee{}).
// 		Where("organization_id = ?", orgID).
// 		Count(&total).
// 		Offset(offset).
// 		Limit(limit).
// 		Find(&employees).Error; err != nil {
// 		return nil, 0, err
// 	}

// 	return employees, total, nil
// }
