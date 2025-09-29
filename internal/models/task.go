package models

import (
	"time"
)

type Task struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Done      bool   `gorm:"default:false"`
	ListID    uint   `gorm:"not null"`
	List      List   `gorm:"foreignKey:ListID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
