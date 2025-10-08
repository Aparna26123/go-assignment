package models

type List struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Position int    `json:"position"`
	BoardID  uint   `json:"board_id"`
}
