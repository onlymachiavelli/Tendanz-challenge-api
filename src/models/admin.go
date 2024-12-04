package models

type Admin struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Email     string `json:"email" gorm:"unique;not null"`
	Identity  string `json:"identity" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Phone     string `json:"phone" gorm:"unique;not null"`
	Verified  bool   `json:"verified" gorm:"default:false"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}