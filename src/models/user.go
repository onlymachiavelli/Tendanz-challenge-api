package models

import "time"

type Client struct {
	ID        uint   `gorm:"autoIncrement;primaryKey" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	Phone string `json:"phone" gorm:"unique"`
	Verified bool   `json:"verified"`
	VerifiedAt time.Time `json:"verified_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}