package service

import (
	"context"

	"github.com/wubshet-kebede/server/internal/domain/model"
	"github.com/wubshet-kebede/server/internal/domain/repo"
)

type HospitalService struct {
    repo repo.HospitalRepository // The Port (Interface)
}

func NewHospitalService(r repo.HospitalRepository) *HospitalService {
    return &HospitalService{repo: r}
}

func (s *HospitalService) RegisterHospital(ctx context.Context, name, address, tinNumber, contactPerson, contactPhone string) (*model.Hospital, error) {
    // 1. Use the Domain Model to create a new hospital (Validation happens here)
    h, err := model.NewHospital(name, address, tinNumber, contactPerson, contactPhone)
    if err != nil {
        return nil, err
    }

    // 2. Use the Repository Adapter to save it to Postgres
    err = s.repo.CreateHospital(ctx, h)
    if err != nil {
        return nil, err
    }

    return h, nil
}