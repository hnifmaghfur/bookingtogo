package repository

import (
	"github.com/hnifmaghfur/bookingtogo/internal/domain"
	"github.com/hnifmaghfur/bookingtogo/internal/interfaces"

	"gorm.io/gorm"
)

type familyListRepository struct {
	db *gorm.DB
}

var _ interfaces.FamilyListRepository = (*familyListRepository)(nil)

func (r *familyListRepository) CreateBulkFamilyList(familyLists []domain.FamilyList) error {
	return r.db.Create(&familyLists).Error
}

func NewFamilyListRepository(db *gorm.DB) interfaces.FamilyListRepository {
	return &familyListRepository{db: db}
}

func (r *familyListRepository) CreateFamilyList(familyList *domain.FamilyList) error {
	return r.db.Create(familyList).Error
}

func (r *familyListRepository) GetFamilyListByID(flId uint) (*domain.FamilyList, error) {
	var familyList domain.FamilyList
	err := r.db.First(&familyList, flId).Error
	return &familyList, err
}

func (r *familyListRepository) GetAllFamilyListsByUserID(cstID uint) ([]domain.FamilyList, error) {
	var familyLists []domain.FamilyList
	err := r.db.Debug().Where("cst_id = ?", cstID).Find(&familyLists).Error
	return familyLists, err
}

func (r *familyListRepository) UpdateFamilyList(familyList *domain.FamilyList) error {
	return r.db.Save(familyList).Error // Save akan melakukan update jika ID ada, create jika tidak
}

func (r *familyListRepository) DeleteFamilyList(id uint) error {
	return r.db.Delete(&domain.FamilyList{}, id).Error
}
