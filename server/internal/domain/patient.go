package domain



type Patient struct {
	ID            uint     `gorm:"primaryKey" json:"id"`
	HospitalID    uint     `gorm:"not null" json:"hospital_id"`
	FullName      string   `gorm:"size:255;not null" json:"full_name"`
	CurrentRoomID uint     `json:"current_room_id"`
	Room          Room     `gorm:"foreignKey:CurrentRoomID"`
	MedicalStatus string   `json:"medical_status"` // "stable", "critical"
}