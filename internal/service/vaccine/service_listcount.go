package vaccine

func (s *VaccineService) GetVaccineCount(title string) (int, error) {
	cnt, err := s.VaccineRepo.GetVaccineCount(title)
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}
