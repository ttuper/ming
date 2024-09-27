package anime

import "ming/internal/models"

func (s *AnimeService) GetAnimeByID(id int) (*models.Anime, []*models.Resource, error) {
	anime, err := s.AnimeRepo.GetAnimeByID(id)
	if err != nil {
		return nil, nil, err
	}
	// 获取资源
	resources, err := s.ResourceRepo.GetResourceListByAnimeID(id)
	if err != nil {
		return nil, nil, err
	}
	return anime, resources, nil
}
