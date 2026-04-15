package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/utils"
)

type contextKey string

const (
	UserIDKey         contextKey = "user_id"
	RoleKey           contextKey = "role"
	HospitalIDKey contextKey = "hospital_id"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessCookie, errA := r.Cookie("access_token")
		refreshCookie, errR := r.Cookie("refresh_token")

		if errA != nil && errR != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
		}

		var accessClaims map[string]interface{}
		var accessToken, refreshToken string

		if errA == nil {
			accessToken = accessCookie.Value
		}
		if errR == nil {
			refreshToken = refreshCookie.Value
		}

		claims, err := utils.VerifyJWT(accessToken)
		if err != nil {
			refreshClaims, refreshErr := utils.VerifyJWT(refreshToken)
			if refreshErr != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userID := fmt.Sprintf("%v", refreshClaims["id"])
			role := fmt.Sprintf("%v", refreshClaims["role"])
			HospitalID := fmt.Sprintf("%v", refreshClaims["hospital_id"])

			newAccess, genErr := utils.GenerateJWT(userID, HospitalID, role, 15*time.Minute)
			if genErr != nil {
				http.Error(w, "Failed to refresh token", http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:     "access_token",
				Value:    newAccess,
				Path:     "/",
				HttpOnly: true,
				Secure:   false,
				SameSite: http.SameSiteLaxMode,
				Expires:  time.Now().Add(15 * time.Minute),
			})

			accessClaims = map[string]interface{}{
				"id":              userID,
				"role":            role,
				"hospital_id":   HospitalID,
			}
		} else {
			accessClaims = claims
		}

		userID := fmt.Sprintf("%v", accessClaims["id"])
		role := fmt.Sprintf("%v", accessClaims["role"])
		HospitalID := fmt.Sprintf("%v", accessClaims["hospital_id"])

		hosUUID, hosUuuidError := uuid.Parse(HospitalID)
		userUUID, userUuidError := uuid.Parse(userID)
		if hosUuuidError != nil || userUuidError != nil {

			http.Error(w, "internal error", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userUUID)
		ctx = context.WithValue(ctx, HospitalIDKey, hosUUID)
		ctx = context.WithValue(ctx, RoleKey, role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}