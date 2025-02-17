package models

type Event struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:200;not null"`
	Description string `gorm:"type:text"`
	Date        string `gorm:"not null"`
	Location    string `gorm:"size:255"`
	CreatedBy   uint
}