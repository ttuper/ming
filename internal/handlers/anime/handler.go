package anime

import anime_service "ming/internal/service/anime"

type AnimeHandler struct {
	AnimeService *anime_service.AnimeService
}

func NewAnimeHandler(as *anime_service.AnimeService) *AnimeHandler {
	return &AnimeHandler{AnimeService: as}
}