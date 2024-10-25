package models

type Vaccine struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"fullName"`
	Category    string `json:"category"`
	Doses       string `json:"doses"`
	Description string `json:"description"`
	IsFree      int    `json:"isFree"`
}
