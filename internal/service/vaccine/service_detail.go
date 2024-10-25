package vaccine

import "ming/internal/models"

func (s *VaccineService) GetVaccineByID(id int) (*models.Vaccine, []*models.VaccineDetail, error) {
	vaccine, err := s.VaccineRepo.GetVaccineByID(id)
	if err != nil {
		return nil, nil, err
	}
	vaccineDetails, err := s.VaccineDetailRepo.GetVaccineDetailsByVaccineID(id)
	return vaccine, vaccineDetails, nil
}
