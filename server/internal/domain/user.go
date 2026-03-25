package domain

type User struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	FullName   string   `gorm:"size:255;not null" json:"full_name"`
	Email      string   `gorm:"uniqueIndex;not null" json:"email"`
	Password   string   `json:"-"` // Hashed password, never exported to JSON
	Role       string   `gorm:"size:50;not null" json:"role"` // "doctor", "staff", "admin"
	HospitalID uint     `gorm:"not null" json:"hospital_id"`
	Hospital   Hospital `gorm:"foreignKey:HospitalID" json:"hospital"`
}
