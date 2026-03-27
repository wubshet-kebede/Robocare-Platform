package auth

import (
	"errors"
	"time"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/features"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/utils"
)

func Login(input model.UserLoginInput) (map[string]string, error) {
	fetchUser, err := features.GetUserByEmailOrPhone(input.Identifier)
	if err != nil {
		return nil, err
	}

	if err := utils.ComparePassword(fetchUser.PasswordHash, input.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	accessToken, err := utils.GenerateJWT(fetchUser.ID.String(), fetchUser.HospitalID.String(), string(fetchUser.Role), 15*time.Minute)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateJWT(fetchUser.ID.String(), fetchUser.HospitalID.String(), string(fetchUser.Role), 7*24*time.Hour)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}
