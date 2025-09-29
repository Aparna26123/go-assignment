package models

type Board struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	UserID uint
}
