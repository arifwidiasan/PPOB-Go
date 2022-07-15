package model

import "time"

type Payment_method struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CodeBank  string    `gorm:"unique;not null" json:"code_bank"`
	Name      string    `gorm:"not null" json:"name"`
	Status    *bool     `gorm:"default:true" json:"status"`
	CodeVA    string    `gorm:"not null" json:"code_va"`
	Range     uint64    `gorm:"default:1;not null" json:"range"`
	CreatedAt time.Time `json:"created_at"`
}
