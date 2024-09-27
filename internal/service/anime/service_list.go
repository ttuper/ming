package anime

import "ming/internal/models"

// 获取动漫列表，分页
func (s *AnimeService) GetAnimeList(page int, pageSize int, title string) ([]*models.Anime, error) {
	// 获取动漫列表
	animes, err := s.AnimeRepo.GetAnimeList(page, pageSize, title)
	if err != nil {
		return nil, err
	}
	return animes, nil
}
