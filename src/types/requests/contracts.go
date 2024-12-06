package requests

import "time"

type CREATELIFECONTRACT struct {
	PolicyType                        string  `json:"policy_type" binding:"required"`
	FaceAmount                        int     `json:"face_amount" binding:"required"`
	PremiumMode                       string  `json:"premium_mode" binding:"required"`
	PremiumAmount                     float64 `json:"premium_amount" binding:"required"`
	PolicyTerm                        int     `json:"policy_term" binding:"required"`
	BenificiaryName                   string  `json:"benificiary_name" binding:"required"`
	BenificiaryRelationship           string  `json:"benificiary_relationship" binding:"required"`
	ContingentBenificiaryName         string  `json:"contingent_benificiary_name" binding:"required"`
	ContingentBenificiaryRelationship string  `json:"contingent_benificiary_relationship" binding:"required"`
	EffectiveDate time.Time `json:"effective_date" binding:"required"`
	ExpirationDate time.Time `json:"expiration_date" binding:"required"`
}

type UpdateLifeContract struct {	
	FaceAmount                        int     `json:"face_amount" binding:"required"`
	PremiumAmount                     float64 `json:"premium_amount" binding:"required"`
	PolicyTerm                        int     `json:"policy_term" binding:"required"`
	BenificiaryName                   string  `json:"benificiary_name" binding:"required"`
	EffectiveDate time.Time `json:"effective_date" binding:"required"`
	ExpirationDate time.Time `json:"expiration_date" binding:"required"`
}

type AcceptRejectLifeContract struct {
	Message string `json:"message" binding:"required"`
}