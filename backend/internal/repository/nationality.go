package repository

import (
	"github.com/hnifmaghfur/bookingtogo/internal/domain"
	"github.com/hnifmaghfur/bookingtogo/internal/interfaces"

	"gorm.io/gorm"
)

type nationalityRepository struct {
	db *gorm.DB
}

// Pastikan nationalityRepository mengimplementasikan service.NationalityRepository.
// Interface service.NationalityRepository akan diperbarui hanya untuk GetAllNationalities.
var _ interfaces.NationalityRepository = (*nationalityRepository)(nil)

func NewNationalityRepository(db *gorm.DB) interfaces.NationalityRepository {
	return &nationalityRepository{db: db}
}

// GetAllNationalities hanya untuk mendapatkan semua daftar kebangsaan yang ada.
func (r *nationalityRepository) GetAllNationalities() ([]domain.Nationality, error) {
	var nationalities []domain.Nationality
	err := r.db.Find(&nationalities).Error
	return nationalities, err
}
