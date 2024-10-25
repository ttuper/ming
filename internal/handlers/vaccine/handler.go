package vaccine

import vaccine_service "ming/internal/service/vaccine"

type VaccineHandler struct {
	VaccineService *vaccine_service.VaccineService
}

func NewVaccineHandler(s *vaccine_service.VaccineService) *VaccineHandler {
	return &VaccineHandler{VaccineService: s}
}
