package model

import "time"

type Virtual_account struct {
	ID              uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	VAID            string         `gorm:"unique;not null" json:"va_id"`
	ExternalID      string         `gorm:"not null" json:"external_id"`
	VANumber        string         `json:"va_number"`
	Price           uint           `gorm:"default:1" json:"price"`
	CreatedAt       time.Time      `json:"created_at"`
	UserID          uint           `json:"user_id"`
	User            User           `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PaymentMethodID uint           `json:"payment_method_id"`
	Payment_Method  Payment_method `json:"payment_method" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
