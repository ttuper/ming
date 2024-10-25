package repository

import (
	"ming/internal/db"
	"ming/internal/models"
	"ming/pkg/logger"
)

type VaccineDetailRepository struct{}

func NewVaccineDetailRepository() *VaccineDetailRepository {
	return &VaccineDetailRepository{}
}
func (r *VaccineDetailRepository) GetVaccineDetailsByVaccineID(vaccineId int) ([]*models.VaccineDetail, error) {
	var vaccineDetails []*models.VaccineDetail
	query := db.DB.Model(&models.VaccineDetail{})

	query = query.Where("vaccine_id = ?", vaccineId)

	err := query.Find(&vaccineDetails).Error
	if err != nil {
		logger.Error("Failed to get vaccine list")
		return nil, err
	}
	return vaccineDetails, nil
}
