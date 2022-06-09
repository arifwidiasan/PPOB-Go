package model

import "time"

type Product struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CodeProduct   string    `gorm:"unique;not null" json:"code_product"`
	Name          string    `gorm:"unique;not null" json:"name"`
	Status        *bool     `gorm:"default:true" json:"status"`
	Price         int       `gorm:"default:1;not null" json:"price"`
	Qty           int       `gorm:"default:1;not null" json:"qty"`
	CreatedAt     time.Time `json:"created_at"`
	ProductTypeID uint      `json:"product_type_id"`
	OperatorID    uint      `json:"operator_id"`
}

type ProductResponse struct {
	ID              uint      `json:"id"`
	CodeProduct     string    `json:"code_product"`
	Name            string    `json:"name"`
	Status          bool      `json:"status"`
	Price           int       `json:"price"`
	Qty             int       `json:"qty"`
	CreatedAt       time.Time `json:"created_at"`
	ProductTypeID   uint      `json:"product_type_id"`
	OperatorID      uint      `json:"operator_id"`
	ProductTypeName string    `json:"product_type_name"`
	OperatorName    string    `json:"operator_name"`
}
