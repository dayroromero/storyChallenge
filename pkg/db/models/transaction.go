package models

import (
	"time"
)

type Transaction struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	AccountID uint `gorm:"foreignKey:ID"`
	Date      time.Time
	Amount    float64
	Type      string
	CreatedAt time.Time `gorm:"default:current_timestamp"`
}
