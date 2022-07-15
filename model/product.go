package model

import "time"

type Product struct {
	ID            uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	CodeProduct   string       `gorm:"unique;not null" json:"code_product"`
	Name          string       `gorm:"unique;not null" json:"name"`
	Status        *bool        `gorm:"default:true" json:"status"`
	Price         uint         `gorm:"default:1;not null" json:"price"`
	Qty           uint         `gorm:"default:1;not null" json:"qty"`
	CreatedAt     time.Time    `json:"created_at"`
	ProductTypeID uint         `json:"product_type_id"`
	Product_type  Product_type `json:"product_type" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OperatorID    uint         `json:"operator_id"`
	Operator      Operator     `json:"operator" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
