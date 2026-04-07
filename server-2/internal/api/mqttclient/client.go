package mqttclient

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)
type MQTTClientManager struct {
	client mqtt.Client
}

type Config struct {
	Broker   string 
	Port     int    
	Username string 
	Password string 
	ClientID string 
}


func NewClient(cfg Config) (*MQTTClientManager, error) {
	opts := mqtt.NewClientOptions()
	brokerURL := fmt.Sprintf("%s:%d", cfg.Broker, cfg.Port)
	opts.AddBroker(brokerURL)
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)
	opts.SetClientID(cfg.ClientID)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(5 * time.Second)

	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("failed to connect to MQTT broker: %w", token.Error())
	}
    fmt.Println("MQTT connected ")
	return &MQTTClientManager{
		client: client,
	}, nil
}


func (m *MQTTClientManager) Close() {
	if m.client != nil && m.client.IsConnected() {
		m.client.Disconnect(250)
	}
}