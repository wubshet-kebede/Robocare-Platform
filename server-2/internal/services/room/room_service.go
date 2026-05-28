package room

import (
	"errors"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/room"
)
func RoomService(input model.Room) (*model.Room, error) {
   
    if input.Capacity <= 0 {
        input.Capacity = 1
    }
    if input.RoomNumber == "" {
    return nil, errors.New("room number required")
}

if input.Floor < 0 {
    return nil, errors.New("invalid floor")
}

    return room.RoomRepository(input)
}