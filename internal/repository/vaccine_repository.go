package repository

import (
	"ming/internal/db"
	"ming/internal/models"
	"ming/pkg/logger"
)

type VaccineRepository struct{}

func NewVaccineRepository() *VaccineRepository {
	return &VaccineRepository{}
}

func (r *VaccineRepository) CreateVaccine(vaccine *models.Vaccine) error {
	err := db.DB.Create(vaccine).Error
	if err != nil {
		logger.Error("Failed to create anime")
		return err
	}
	return nil
}

func (r *VaccineRepository) GetVaccineByID(id int) (*models.Vaccine, error) {
	var vaccine models.Vaccine
	err := db.DB.First(&vaccine, id).Error
	if err != nil {
		logger.Error("Failed to get anime by ID")
		return nil, err
	}
	return &vaccine, nil
}

func (r *VaccineRepository) GetVaccineList(keyword string) ([]*models.Vaccine, error) {
	var vaccines []*models.Vaccine
	query := db.DB.Model(&models.Vaccine{})

	if keyword != "" {
		query = query.Where("name LIKE ? OR full_name LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	err := query.Find(&vaccines).Error
	if err != nil {
		logger.Error("Failed to get vaccine list")
		return nil, err
	}
	return vaccines, nil
}

func (r *VaccineRepository) GetVaccineCount(title string) (int64, error) {
	query := db.DB.Model(&models.Vaccine{})
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
