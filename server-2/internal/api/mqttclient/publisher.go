package mqttclient

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
)


func NewMQTTPublisher(broker string, port int, secrets map[string]string) *MQTTPublisher {
	return &MQTTPublisher{
		clients: make(map[string]*MQTTClientManager),
		secrets: secrets,
		broker:  broker,
		port:    port,
	}
}


func (p *MQTTPublisher) PublishNavGoal(hospitalUUID, robotUUID string, goal GoalPayload) error {
	p.mu.Lock()
	defer p.mu.Unlock()

    
	client, ok := p.clients[hospitalUUID]
	if !ok {
		_, exists := p.secrets[hospitalUUID]
		if !exists {
			return fmt.Errorf("no secret found for hospital %s", hospitalUUID)
		}
       brokerUser := os.Getenv("MQTT_USERNAME")
       brokerPass := os.Getenv("MQTT_PASSWORD")
		newClient, err := NewClient(Config{
			Broker:   p.broker,
			Port:     p.port,
			Username: brokerUser,
			Password: brokerPass,
			ClientID: fmt.Sprintf("backend-%s", hospitalUUID),
		})
		if err != nil {
			return fmt.Errorf("failed to create MQTT client: %w", err)
		}
		p.clients[hospitalUUID] = newClient
		client = newClient
	}

	topic := buildTopic(hospitalUUID, robotUUID)

	payloadBytes, err := json.Marshal(goal)
	log.Printf("📡 Publishing to topic: %s", topic)
    log.Printf("📦 Payload: %s", string(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to marshal goal payload: %w", err)
	}
	fmt.Println("Using secret:", p.secrets[hospitalUUID])

	hmacSignature := signPayload(payloadBytes, p.secrets[hospitalUUID])
	goal.HMAC = hmacSignature
	payloadBytes, err = json.Marshal(goal)
	if err != nil {
		return fmt.Errorf("failed to marshal goal payload with HMAC: %w", err)
	}
    if !client.client.IsConnected() {
	return fmt.Errorf("client not connected")
}
	token := client.client.Publish(topic, 1, false, payloadBytes) // QoS=1
	token.Wait()
	if token.Error() != nil {
		return fmt.Errorf("failed to publish nav goal: %w", token.Error())
	}

	return nil
}

func buildTopic(hospitalUUID, robotUUID string) string {
	return fmt.Sprintf("hospital/%s/robot/%s/goal", hospitalUUID, robotUUID)
}


func signPayload(payload []byte, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(payload)
	return hex.EncodeToString(h.Sum(nil))
}

