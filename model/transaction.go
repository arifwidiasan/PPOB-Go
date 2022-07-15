package model

import "time"

type Transaction struct {
	ID              uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CodeTransaction string         `gorm:"unique;not null" json:"code_transaction"`
	Status          string         `gorm:"default:pending" json:"status"`
	ChargeNumber    string         `gorm:"not null" json:"charge_number"`
	Token           string         `json:"token"`
	Price           uint           `gorm:"default:1;not null" json:"price"`
	CreatedAt       time.Time      `json:"created_at"`
	UserID          uint           `json:"user_id"`
	User            User           `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductID       uint           `json:"product_id"`
	Product         Product        `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PaymentMethodID uint           `json:"payment_method_id"`
	Payment_Method  Payment_method `json:"payment_method" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
