package robot

import (
    "errors"
    "github.com/google/uuid"
    "github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
    "github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/robot"
)

func RegisterRobot(hospitalID uuid.UUID, serial string, modelName string) (*model.Robot, error) {
    // 1. Check if this physical robot is already in the system
    existing, _ := robot.GetBySerial(serial)
    if existing != nil && existing.ID != uuid.Nil {
        return nil, errors.New("serial number already registered")
    }

    // 2. Build the Robot object
    newBot := &model.Robot{
        ID:           uuid.New(),
        HospitalID:   hospitalID,
        SerialNumber: serial,
        Model:        modelName,
        Status:       model.RobotOffline, // Default state
        BatteryLevel: 0,
    }

    // 3. Save to Database
    if err := robot.Create(newBot); err != nil {
        return nil, err
    }

    return newBot, nil
}