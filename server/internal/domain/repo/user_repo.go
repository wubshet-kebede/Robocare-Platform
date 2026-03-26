package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/server/internal/domain/model"
)
type UserRepo interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}