package repo
import (
	"context"
	"github.com/google/uuid"
	"github.com/wubshet-kebede/server/internal/domain/model"
)
type PatientRepo interface {
	CreatePatient(ctx context.Context, patient *model.Patient) error
	GetPatientByID(ctx context.Context, id uuid.UUID) (*model.Patient, error)
	UpdatePatient(ctx context.Context, patient *model.Patient) error
	DeletePatient(ctx context.Context, id uuid.UUID) error
}