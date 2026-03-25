package domain


type Room struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	HospitalID uint     `gorm:"not null" json:"hospital_id"`
	RoomNumber string   `gorm:"size:100;not null" json:"room_number"` // "Bed-01"
	X          float64  `json:"x_coordinate"` // SLAM X
	Y          float64  `json:"y_coordinate"` // SLAM Y
	Yaw        float64  `json:"yaw"`          // Orientation
	Status     string   `gorm:"default:'available'" json:"status"` 
}