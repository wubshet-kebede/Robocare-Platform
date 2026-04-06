package mqtt

import (
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func RegisterSubscribers(client mqtt.Client) {
    topic := "hospitals/+/robots/+/vitals"
    token := client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
    parts := strings.Split(msg.Topic(), "/")
    if len(parts) < 5 { return }

    hospitalID := parts[1]
    robotID := parts[3]
    HandleVitalsSubmission(hospitalID, robotID, msg.Payload())
})

token.Wait()
if token.Error() != nil {
    println("Subscribe error:", token.Error().Error())
}
}