package robot

import (
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
)

func Create(r *model.Robot) error {
    return db.DB.Create(r).Error
}

func GetBySerial(serial string) (*model.Robot, error) {
    var r model.Robot
    err := db.DB.Where("serial_number = ?", serial).First(&r).Error
    return &r, err
}