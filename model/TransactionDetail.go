package model

type TransactionDetail struct {
	Currency    string  `json:"currency"`
	GrossAmount float64 `json:"gross_amount"`
	OrderID     string  `json:"order_id"`
}
