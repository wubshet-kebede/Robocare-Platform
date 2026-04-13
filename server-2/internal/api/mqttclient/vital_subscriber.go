package mqttclient

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)



func (s *VitalsSubscriber) Subscribe() error {

	topic := "hospital/+/robot/+/patient/+/vitals"

	s.client.client.Subscribe(topic, 1, func(c mqtt.Client, msg mqtt.Message) {

		fmt.Println("📥 Topic:", msg.Topic())
		fmt.Println("Payload:", string(msg.Payload()))

		// TODO: decode JSON
		// TODO: save to DB
	})

	return nil
}