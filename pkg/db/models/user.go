package models

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique"`
	Email    string
}
