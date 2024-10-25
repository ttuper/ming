package models

type VaccineDetail struct {
	ID        int    `json:"id"`
	VaccineID int    `json:"vaccine_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}
