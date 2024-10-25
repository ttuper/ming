package vaccine

import "ming/internal/repository"

type VaccineService struct {
	VaccineRepo       *repository.VaccineRepository
	VaccineDetailRepo *repository.VaccineDetailRepository
}

func NewVaccineService() *VaccineService {
	return &VaccineService{
		VaccineRepo:       repository.NewVaccineRepository(),
		VaccineDetailRepo: repository.NewVaccineDetailRepository(),
	}
}
