package models

type Cart struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
}
