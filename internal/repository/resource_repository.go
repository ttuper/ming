package repository

import (
	"ming/internal/db"
	"ming/internal/models"
	"ming/pkg/logger"
)

type ResourceRepository struct{}

func NewResourceRepository() *ResourceRepository {
	return &ResourceRepository{}
}

// 获取资源列表
func (r *ResourceRepository) GetResourceListByAnimeID(animeID int) ([]*models.Resource, error) {
	var resources []*models.Resource
	query := db.DB.Model(&models.Resource{})

	err := query.Where("anime_id = ?", animeID).Find(&resources).Error
	if err != nil {
		logger.Error("Failed to get anime by ID")
		return nil, err
	}
	return resources, nil
}
