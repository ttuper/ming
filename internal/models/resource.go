package models

type Resource struct {
	ID        uint      `json:"id"`
	AnimeID   int32     `json:"anime_id"`
	Type      int8      `json:"type"`
	Url       string    `json:"url"`
	Code      string    `json:"code"`
}
