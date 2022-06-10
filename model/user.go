package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string    `gorm:"unique" json:"email"`
	Username  string    `gorm:"unique" json:"username"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Balance   uint      `gorm:"default:0" json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}
