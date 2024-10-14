package models

type Feedback struct {
	ID   int    `json:"id"`
	Type int    `json:"type"`
	Desc string `json:"desc"`
}
