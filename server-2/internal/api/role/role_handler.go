package role
import (
	"net/http"
	"encoding/json"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/role"
)
func GetRoles(w http.ResponseWriter, r *http.Request) {
    roles, err := role.GetRoles()
    if err != nil {
        http.Error(w, "Failed to fetch roles", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(roles)
}