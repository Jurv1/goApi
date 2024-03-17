package models

type Quest struct {
	ID   int    `json:"Id"`
	Name string `json:"name"`
	Cost int    `json:"cost"`
}
