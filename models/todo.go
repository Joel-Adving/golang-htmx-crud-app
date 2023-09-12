package models

type Todo struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `json:"title"`
	Completed bool   `json:"completed" gorm:"default:false"`
}
