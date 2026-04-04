package department

import (
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/department"
)
func CreateDepartment(dept model.Department) (*model.Department, error) {
	createDept, err := department.CreateDepartment(dept)
	if err != nil {
		return &model.Department{}, err
	}
	return createDept, nil
}