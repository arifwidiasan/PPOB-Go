package model

import "time"

type Admin struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:30" json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminLogin struct {
	Username string `json:"username" example:"arifwidiasan"`
	Password string `json:"password" example:"rahasia87"`
}
