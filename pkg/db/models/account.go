package models

import (
	"time"
)

type Account struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	AccountNumber string `gorm:"unique"`
	Balance       float64
	UserId        uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
