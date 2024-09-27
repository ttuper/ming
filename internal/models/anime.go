package models

type Anime struct {
	ID          int      `json:"id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	Desc        string    `json:"desc"`
	Genres      string    `json:"genres"`
	Poster      string    `json:"poster"`
	ReleaseDate string    `json:"release_date"`
	Score       string   `json:"score"`
}
