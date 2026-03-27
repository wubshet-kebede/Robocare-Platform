package auth
import (
  "github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
  "github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/auth"

)
func Signup(input model.HospitalRequest) (model.Hospital, model.User, error) {
	createdHospital, admin, err := auth.CreateHospitalWithAdmin(input)
	return createdHospital, admin, err
}