package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/server/internal/domain/model"
)
type RobotRepo interface {
	CreateRobot(ctx context.Context, robot *model.Robot) error
	GetRobotByID(ctx context.Context, id uuid.UUID) (*model.Robot, error)
	UpdateRobot(ctx context.Context, robot *model.Robot) error
	DeleteRobot(ctx context.Context, id uuid.UUID) error
}