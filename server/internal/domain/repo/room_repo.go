package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/server/internal/domain/model"
)
type RoomRepo interface {
	CreateRoom(ctx context.Context, room *model.Room) error
	GetRoomByID(ctx context.Context, id uuid.UUID) (*model.Room, error)
	UpdateRoom(ctx context.Context, room *model.Room) error
	DeleteRoom(ctx context.Context, id uuid.UUID) error
}