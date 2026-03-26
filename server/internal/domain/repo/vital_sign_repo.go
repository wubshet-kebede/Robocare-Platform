package repo
import (
	"context"
	"github.com/google/uuid"
	"github.com/wubshet-kebede/server/internal/domain/model"
)
type VitalSignRepo interface {
	CreateVitalSign(ctx context.Context, vitalSign *model.VitalSign) error
	GetVitalSignByID(ctx context.Context, id uuid.UUID) (*model.VitalSign, error)
	UpdateVitalSign(ctx context.Context, vitalSign *model.VitalSign) error
	DeleteVitalSign(ctx context.Context, id uuid.UUID) error
}