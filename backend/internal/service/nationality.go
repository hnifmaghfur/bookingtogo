package service

import (
	"github.com/hnifmaghfur/bookingtogo/internal/domain"
	"github.com/hnifmaghfur/bookingtogo/internal/interfaces"
)

type nationalityService struct {
	nationalityRepo interfaces.NationalityRepository
}

func NewNationalityService(nationalityRepo interfaces.NationalityRepository) interfaces.NationalityService {
	return &nationalityService{nationalityRepo: nationalityRepo}
}

func (s *nationalityService) GetAllNationalities() ([]domain.Nationality, error) {
	return s.nationalityRepo.GetAllNationalities()
}
