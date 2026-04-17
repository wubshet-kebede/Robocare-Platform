package mqttclient

import (
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/api/ws"
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

func InitMQTTSubscriber(broker string, port int,  wsManager *ws.Manager) {
    cfg := Config{
        Broker:   broker,
        Port:     port,
        Username: os.Getenv("MQTT_USERNAME"),
        Password: os.Getenv("MQTT_PASSWORD"),
        ClientID: "backend-subscriber",       
    }

    mqttSub, err := NewClient(cfg)
    if err != nil {
        log.Fatal("failed to init MQTT subscriber:", err)
    }
    token := mqttSub.client.Subscribe("hospital/+/robot/+/patient/+/vitals", 1, func(client mqtt.Client, msg mqtt.Message) {
    HandleVitalsMessage(wsManager, msg)
   })
    token.Wait()
    if token.Error() != nil {
       log.Fatal("failed to subscribe:", token.Error())
    }
    
    log.Println("MQTT subscriber initialized and listening for vitals")
}