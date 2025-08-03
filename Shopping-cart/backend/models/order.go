package models

type Order struct {
	ID     uint `gorm:"primaryKey"`
	CartID uint
	UserID uint
}
