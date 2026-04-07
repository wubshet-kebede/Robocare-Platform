package mqttclient

import (
	"log"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
)

func InitMQTTPublisher(broker string, port int) {
	
	hospitalSecrets := make(map[string]string)
	var hospitals []struct {
		ID     string
		Secret string
	}
	if err := db.DB.Model(&model.Hospital{}).Select("id", "secret").Find(&hospitals).Error; err != nil {
		log.Fatal("failed to load hospital secrets:", err)
	}
	for _, h := range hospitals {
		hospitalSecrets[h.ID] = h.Secret
	}

	mqttPub := NewMQTTPublisher(broker, port, hospitalSecrets)
	SetMQTTPublisher(mqttPub) 
	log.Println("MQTT publisher initialized with all hospitals")
}