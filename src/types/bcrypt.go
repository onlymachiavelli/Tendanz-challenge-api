package types

type Cost struct {
	Cost int `json:"cost" validate:"required"`
}