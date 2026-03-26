package model

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)
type Hospital struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

func NewHospital(name, address string) (*Hospital, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if address == "" {
		return nil, errors.New("address is required")
	}
    slug := strings.ToLower(strings.ReplaceAll(name, " ", "-"))
	return &Hospital{
		ID:        uuid.New(),
		Name:      name,
		Slug:      slug,
		Address:   address,
		CreatedAt: time.Now(),
	}, nil
}