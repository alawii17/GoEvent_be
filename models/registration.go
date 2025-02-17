package models

type Registration struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint
	EventID uint
}