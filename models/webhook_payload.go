package models

type WebhookPayload struct {
	ID          string `json:"id" bson:"id"`
	Amount      int    `json:"amount" bson:"amount"`
	Currency    string `json:"currency" bson:"currency"`
	CreatedAt   int64  `json:"created_at_time" bson:"created_at_time"`
	Timestamp   int64  `json:"timestamp" bson:"timestamp"`
	Cause       string `json:"cause" bson:"cause"`
	FullName    string `json:"full_name" bson:"full_name"`
	AccountName string `json:"account_name" bson:"account_name"`
	InvoiceURL  string `json:"invoice_url" bson:"invoice_url"`
}
