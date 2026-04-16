package role

import (
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/role"
)
func GetRoles()([]string, error){
	return role.GetRoles()
}