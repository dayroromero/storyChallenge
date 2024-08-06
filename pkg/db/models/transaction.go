package models

import (
	"time"
)

type Transaction struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	AccountID uint
	Amount    float64
	Type      string
	CreatedAt time.Time
}
