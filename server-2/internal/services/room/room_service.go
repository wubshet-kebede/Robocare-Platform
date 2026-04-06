package room

import (
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/room"
)
func RoomService(input model.Room) (*model.Room, error) {
   
    if input.Capacity <= 0 {
        input.Capacity = 1
    }

    return room.RoomRepository(input)
}