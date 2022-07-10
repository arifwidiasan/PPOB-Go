package model

type Callback_payment struct {
	ID         string `gorm:"primaryKey" json:"id"`
	VAID       string `json:"callback_virtual_account_id"`
	ExternalID string `json:"external_id"`
	PaymentID  string `json:"payment_id"`
	VANumber   string `json:"account_number"`
	Name       string `json:"name"`
	Status     string `json:"status"`
}
