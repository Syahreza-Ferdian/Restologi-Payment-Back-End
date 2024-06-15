package model

type CreditCardRequest struct {
	Authentication string `json:"authentication"`
	SaveCard       bool   `json:"save_card"`
	Secure         bool   `json:"secure"`
}
