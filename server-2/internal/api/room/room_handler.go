package room
import (
	"net/http"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/middleware"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/room"
)

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
    role := r.Context().Value(middleware.RoleKey).(string)
    if role != "admin" {
        http.Error(w, "Forbidden: Only admins can manage infrastructure", http.StatusForbidden)
        return
    }

    hospitalID := r.Context().Value(middleware.HospitalIDKey).(uuid.UUID)

    var input model.Room
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    input.HospitalID = hospitalID
    
    room, err := room.RoomService(input)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(room)
}