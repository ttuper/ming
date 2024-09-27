package repository

import (
	"ming/internal/db"
	"ming/internal/models"
	"ming/pkg/logger"
)

type AnimeRepository struct{}

func NewAnimeRepository() *AnimeRepository {
	return &AnimeRepository{}
}

func (r *AnimeRepository) CreateAnime(anime *models.Anime) error {
	err := db.DB.Create(anime).Error
	if err != nil {
		logger.Error("Failed to create anime")
		return err
	}
	return nil
}

func (r *AnimeRepository) GetAnimeByID(id int) (*models.Anime, error) {
	var anime models.Anime
	err := db.DB.First(&anime, id).Error
	if err != nil {
		logger.Error("Failed to get anime by ID")
		return nil, err
	}
	return &anime, nil
}

// 获取动漫列表
func (r *AnimeRepository) GetAnimeList(page int, pageSize int, title string) ([]*models.Anime, error) {
	var animes []*models.Anime
	query := db.DB.Model(&models.Anime{})

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}

	offset := (page - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	err := query.Find(&animes).Error
	if err != nil {
		logger.Error("Failed to get anime by ID")
		return nil, err
	}
	return animes, nil
}

// 获取动漫数量
func (r *AnimeRepository) GetAnimeCount(title string) (int64, error) {
	query := db.DB.Model(&models.Anime{})
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	var count int64
	err := query.Count(&count).Error
	if err != nil {
		logger.Error("Failed to get anime count")
		return 0, err
	}
	return count, nil
}
