package anime


// 获取动漫数量
func (s *AnimeService) GetAnimeCount(title string) (int, error) {
	cnt, err := s.AnimeRepo.GetAnimeCount(title)
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}
