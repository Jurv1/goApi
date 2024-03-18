package models

type Quest struct {
	ID   int64  `json:"Id"`
	Name string `json:"name"`
	Cost int    `json:"cost"`
}
