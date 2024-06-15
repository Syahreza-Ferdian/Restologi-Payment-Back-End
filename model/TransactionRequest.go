package model

type TransactionRequest struct {
	CreditCard         CreditCardRequest `json:"credit_card"`
	CustomerDetails    CustomerDetail    `json:"customer_details"`
	ItemDetails        []ItemDetail      `json:"item_details"`
	TransactionDetails TransactionDetail `json:"transaction_details"`
}
