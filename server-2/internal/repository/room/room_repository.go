package room

import (
	"fmt"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
)

func RoomRepository(room model.Room) (*model.Room, error) {
    var existing model.Room
    err := db.DB.Where("hospital_id = ? AND room_number = ?", room.HospitalID, room.RoomNumber).First(&existing).Error
    if err == nil {
        return nil, fmt.Errorf("room number %s already exists in this hospital", room.RoomNumber)
    }
    if err := db.DB.Create(&room).Error; err != nil {
        return nil, err
    }
    return &room, nil
}