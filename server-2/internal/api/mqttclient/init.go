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

func NewVitalsSubscriber(cfg Config) (*VitalsSubscriber, error) {

	client, err := NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &VitalsSubscriber{
		client: client,
	}, nil
}
// func InitMQTTSubscriber(broker string, port int, hospitalSecrets map[string]string) {
//     cfg := Config{
//         Broker:   broker,
//         Port:     port,
//         Username: os.Getenv("MQTT_USERNAME"), // only if broker requires auth
//         Password: os.Getenv("MQTT_PASSWORD"),
//         ClientID: "backend-subscriber",       // must be unique
//     }

//     mqttSub, err := NewClient(cfg)
//     if err != nil {
//         log.Fatal("failed to init MQTT subscriber:", err)
//     }
//     err = mqttSub.Subscribe("hospital/+/patient/+/vitals", func(client mqtt.Client, msg mqtt.Message) {
//         HandleVitalsMessage(msg, hospitalSecrets)
//     })
//     if err != nil {
//         log.Fatal("failed to subscribe:", err)
//     }

//     log.Println("MQTT subscriber initialized and listening for vitals")
// }