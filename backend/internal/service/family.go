package service

import (
	"errors"

	"github.com/hnifmaghfur/bookingtogo/internal/domain"
	"github.com/hnifmaghfur/bookingtogo/internal/interfaces"
)

type familyListService struct {
	familyListRepo interfaces.FamilyListRepository
	customerRepo   interfaces.CustomerRepository // Perlu customerRepo untuk validasi CustomerID
}

func (s *familyListService) CreateBulkFamilyList(familyLists []domain.FamilyList) error {
	if len(familyLists) == 0 {
		return errors.New("family list array cannot be empty")
	}
	cstID := familyLists[0].CstID
	for _, f := range familyLists {
		if f.CstID != cstID {
			return errors.New("all family members must have the same cst_id")
		}
		if f.FlName == "" || f.FlRelation == "" {
			return errors.New("nama dan relasi tidak boleh kosong")
		}
	}
	// Validasi customer ID
	_, err := s.customerRepo.GetCustomerByID(cstID)
	if err != nil {
		return errors.New("customer ID tidak valid: " + err.Error())
	}
	return s.familyListRepo.CreateBulkFamilyList(familyLists)
}

func NewFamilyListService(familyListRepo interfaces.FamilyListRepository, customerRepo interfaces.CustomerRepository) interfaces.FamilyListService {
	return &familyListService{
		familyListRepo: familyListRepo,
		customerRepo:   customerRepo,
	}
}

func (s *familyListService) CreateFamilyList(familyList *domain.FamilyList) error {
	if familyList.FlName == "" || familyList.FlRelation == "" || familyList.CstID == 0 {
		return errors.New("nama, relasi, dan customer ID tidak boleh kosong")
	}
	// Pastikan CustomerID yang diberikan ada
	_, err := s.customerRepo.GetCustomerByID(familyList.CstID)
	if err != nil {
		return errors.New("customer ID tidak valid: " + err.Error())
	}
	return s.familyListRepo.CreateFamilyList(familyList)
}

func (s *familyListService) GetFamilyListByID(id uint) (*domain.FamilyList, error) {
	return s.familyListRepo.GetFamilyListByID(id)
}

func (s *familyListService) GetAllFamilyListsByUserID(userID uint) ([]domain.FamilyList, error) {
	// Pastikan UserID yang diberikan ada
	_, err := s.customerRepo.GetCustomerByID(userID)
	if err != nil {
		return nil, errors.New("customer ID tidak valid: " + err.Error())
	}
	return s.familyListRepo.GetAllFamilyListsByUserID(userID)
}

func (s *familyListService) UpdateFamilyList(familyList *domain.FamilyList) error {
	if familyList.FlID == 0 || familyList.FlName == "" || familyList.FlRelation == "" || familyList.CstID == 0 {
		return errors.New("ID, nama, relasi, dan user ID tidak boleh kosong")
	}
	// Pastikan UserID yang diberikan ada
	_, err := s.customerRepo.GetCustomerByID(familyList.CstID)
	if err != nil {
		return errors.New("customer ID tidak valid: " + err.Error())
	}
	return s.familyListRepo.UpdateFamilyList(familyList)
}

func (s *familyListService) DeleteFamilyList(id uint) error {
	return s.familyListRepo.DeleteFamilyList(id)
}
