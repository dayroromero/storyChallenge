package models

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	UserName string `gorm:"unique"`
	Email    string
}
