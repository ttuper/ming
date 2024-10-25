package vaccine

import "ming/internal/models"

func (s *VaccineService) GetVaccineList(title string) ([]*models.Vaccine, error) {
	vaccines, err := s.VaccineRepo.GetVaccineList(title)
	if err != nil {
		return nil, err
	}
	return vaccines, nil
}
