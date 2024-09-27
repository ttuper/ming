package anime

import "ming/internal/repository"

type AnimeService struct {
	AnimeRepo *repository.AnimeRepository
	ResourceRepo *repository.ResourceRepository
}

func NewAnimeService() *AnimeService {
	return &AnimeService{
		AnimeRepo: repository.NewAnimeRepository(),
		ResourceRepo: repository.NewResourceRepository(),
	}
}