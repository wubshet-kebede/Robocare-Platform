package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)
type RobotStatus string

const (
    RobotIdle       RobotStatus = "idle"
    RobotNavigating RobotStatus = "navigating" // Moving to a patient
    RobotCharging   RobotStatus = "charging"
    RobotEmergency  RobotStatus = "emergency_stop"
    RobotOffline    RobotStatus = "offline"
)
type Robot struct {
    ID            uuid.UUID 
    HospitalID    uuid.UUID 
    
    // Navigation State
    CurrentRoomID *uuid.UUID // Where it is now
    TargetRoomID  *uuid.UUID // Where the doctor sent it
    Status        RobotStatus     // "idle", "charging", "moving", "emergency_stop", "navigating"
    
    // Telemetry from ROS2
    BatteryLevel  int      
    WiFiStrength  int        // Important for video quality in telepresence
    IPAddress     string     // Needed for WebRTC handshake
    
    // Hardware Health
    LastSeen      time.Time 
}

func NewRobot(hospitalID uuid.UUID) (*Robot, error) {
    if hospitalID == uuid.Nil {
        return nil, errors.New("hospital ID is required")
    }
    
    return &Robot{
        ID:           uuid.New(),
        HospitalID:   hospitalID,
        Status:       RobotIdle,
        BatteryLevel: 100,
        LastSeen:     time.Now(),
    }, nil
}
