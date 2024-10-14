package models

type Anime struct {
	ID          int      `json:"id"`
	Title       string    `json:"title"`
	Poster      string    `json:"poster"`
	Desc        string    `json:"desc"`
	Genres      string    `json:"genres"`
	Score       string   `json:"score"`
	ReleaseDate string    `json:"release_date"`
}
