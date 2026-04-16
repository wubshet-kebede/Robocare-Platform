package role

import (
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
)

func  GetRoles() ([]string, error) {
    var roles []string
    err := db.DB.Raw(`
        SELECT e.enumlabel
        FROM pg_enum e
        JOIN pg_type t ON e.enumtypid = t.oid
        WHERE t.typname = 'role_enum';
    `).Scan(&roles).Error
    return roles, err
}