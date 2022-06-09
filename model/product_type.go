package model

import "time"

type Product_type struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Products  []Product `gorm:"ForeignKey:ProductTypeID"`
}
