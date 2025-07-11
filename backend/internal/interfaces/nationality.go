package interfaces

import "github.com/hnifmaghfur/bookingtogo/internal/domain"

type NationalityRepository interface {
	GetAllNationalities() ([]domain.Nationality, error)
}

type NationalityService interface {
	GetAllNationalities() ([]domain.Nationality, error)
}
