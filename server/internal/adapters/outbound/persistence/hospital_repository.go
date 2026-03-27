package persistence

import (
	"context"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/server/internal/domain/model"
	"github.com/wubshet-kebede/server/internal/domain/repo"
	"gorm.io/gorm"
)
type HospitalRepository struct {
	db *gorm.DB
}

// NewHospitalRepository creates a new instance that matches the repo.HospitalRepository interface
func NewHospitalRepository(db *gorm.DB) repo.HospitalRepository {
	return &HospitalRepository{db: db}
}

func (r *HospitalRepository) CreateHospital(ctx context.Context, h *model.Hospital) error {
	return r.db.WithContext(ctx).Create(h).Error
}

func (r *HospitalRepository) GetHospitalBySlug(ctx context.Context, slug string) (*model.Hospital, error) {
	var h model.Hospital
	err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&h).Error
	return &h, err
}
func (r *HospitalRepository) GetHospitalByID(ctx context.Context, id uuid.UUID) (*model.Hospital, error) {
	var h model.Hospital
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&h).Error
	return &h, err
}
func (r *HospitalRepository) UpdateHospital(ctx context.Context, h *model.Hospital) error {
	return r.db.WithContext(ctx).Save(h).Error

}
func (r *HospitalRepository) DeleteHospital(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Hospital{}).Error
}