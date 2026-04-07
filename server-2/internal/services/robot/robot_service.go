package robot

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/robot"
)

func RegisterRobot(hospitalID uuid.UUID, serial string, modelName string) (*model.Robot, error) {
	serial = strings.TrimSpace(serial)
	modelName = strings.TrimSpace(modelName)

	if serial == "" {
		return nil, errors.New("serial number is required")
	}
	existing, err := robot.GetBySerial(serial)
	if err != nil {
		return nil, err
	}
	if existing != nil && existing.ID != uuid.Nil {
		return nil, errors.New("serial number already registered")
	}

	newRobot := &model.Robot{
		HospitalID:   hospitalID,
		SerialNumber: serial,
		Model:        modelName,
		Status:       model.RobotOffline,
		BatteryLevel: 0,
		LastSeen:     time.Now(),
	}

	if err := robot.Create(newRobot); err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return nil, errors.New("serial number already registered")
		}
		return nil, err
	}

	return newRobot, nil
}