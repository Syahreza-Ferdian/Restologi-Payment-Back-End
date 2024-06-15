package model

type TransactionResponse struct {
	Token       string `json:"token"`
	OrderID     string `json:"order_id"`
	GrossAmount int64  `json:"gross_amount"`
	SnapUrl     string `json:"snap_url"`
}
