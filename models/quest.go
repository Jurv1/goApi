package models

type Quest struct {
	ID   int    `json:"id" gorm:"AUTO_INCREMENT; PRIMARY_KEY"`
	Name string `json:"name"`
	Cost int    `json:"cost"`
}
