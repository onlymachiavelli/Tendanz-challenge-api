package models

import "time"


type LifeInsurance struct {

	ID            int     `gorm:"primaryKey;autoIncrement;not null;unique" json:"id"`
	ClientID      int     `json:"client_id" gorm:"not null unique;foreignKey:ClientID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PolicyType    string  `json:"policy_type" gorm:"not null"`
	FaceAmount    int     `json:"face_amount" gorm:"not null"`
	PremiumMode   string  `json:"premium_mode" gorm:"not null"`
	PremiumAmount float64 `json:"premium_amount" gorm:"not null"`
	PolicyTerm    int     `json:"policy_term" gorm:"not null"`

	BenificiaryName string `json:"benificiary_name" gorm:"not null"`
	BenificiaryRelationship string `json:"benificiary_relationship" gorm:"not null"`
	ContingentBenificiaryName string `json:"contingent_benificiary_name" gorm:"not null"`
	ContingentBenificiaryRelationship string `json:"contingent_benificiary_relationship" gorm:"not null"`

	EffectiveDate  time.Time `json:"effective_date" gorm:"not null"`
	ExpirationDate time.Time `json:"expiration_date" gorm:"not null"`


	Status string `json:"status" gorm:"not null"`	//pending or approved or rejected
	Message string `json:"message" gorm:"not null"`	//additional message for the status

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}
