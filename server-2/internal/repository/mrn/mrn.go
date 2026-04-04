package mrn

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"gorm.io/gorm"
)
func GenerateMRN(hospitalID uuid.UUID) (string, error) {
	currentYear := time.Now().Year()

	var seq model.Sequence

	
	tx := db.DB.Begin()

	err := tx.
		Where("hospital_id = ? AND type = ? AND year = ?", hospitalID, "admission", currentYear).
		First(&seq).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			seq = model.Sequence{
				HospitalID: hospitalID,
				Type:       "admission",
				Year:       currentYear,
				LastNumber: 1,
			}
			if err := tx.Create(&seq).Error; err != nil {
				tx.Rollback()
				return "", err
			}
		} else {
			tx.Rollback()
			return "", err
		}
	} else {
		seq.LastNumber++
		if err := tx.Save(&seq).Error; err != nil {
			tx.Rollback()
			return "", err
		}
	}

	tx.Commit()
	mrn := fmt.Sprintf("ADM-%d-%06d", currentYear, seq.LastNumber)

	return mrn, nil
}