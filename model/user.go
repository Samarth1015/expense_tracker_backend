package model

import (
	"time"
)

// type User struct {
// 	ID        string `gorm:"primaryKey"`
// 	UserName  string `gorm:"type:varchar(100);not null"`
// 	Email     string `gorm:"type:varchar(100);unique;not null"`
// 	Role      string `gorm:"type:user_role;default:'user';not null"`
// 	CreatedAt time.Time
// }

type User struct {
	ID        string `gorm:"primaryKey"`
	UserName  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);unique;not null"`
	Role      string `gorm:"type:user_role;default:'user';not null"`
	CreatedAt time.Time
	Expenses  []Expense `gorm:"foreignKey:UserID"`
}
type Expense struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Date        time.Time `gorm:"not null" json:"date"`
	Amount      float64   `gorm:"not null" json:"amount"`
	Description string    `gorm:"type:text" json:"description"`
	Category    string    `gorm:"type:varchar(100);not null" json:"category"`
	UserID      string    `gorm:"not null" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID;references:ID" json:"-"`
}
