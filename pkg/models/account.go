package models

import (
	"time"
)

type Account struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	AccountNumber string `gorm:"unique"`
	UserId        uint   `gorm:"foreignKey:ID"`
	Balance       float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
