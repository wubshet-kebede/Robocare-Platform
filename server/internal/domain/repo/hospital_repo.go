package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/server/internal/domain/model"
)
type HospitalRepo interface {
	CreateHospital(ctx context.Context, hospital *model.Hospital) error
	GetHospitalByID(ctx context.Context, id uuid.UUID) (*model.Hospital, error)
	GetHospitalBySlug(ctx context.Context, slug string) (*model.Hospital, error)
	UpdateHospital(ctx context.Context, hospital *model.Hospital) error
	DeleteHospital(ctx context.Context, id uuid.UUID) error
}