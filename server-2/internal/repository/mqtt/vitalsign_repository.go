package mqtt

import (
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
)
func SaveVitalSign(vitals model.VitalSign) error {
    
    if err := db.DB.Create(&vitals).Error; err != nil {
        return err
    }
    return nil
}