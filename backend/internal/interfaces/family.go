package interfaces

import "github.com/hnifmaghfur/bookingtogo/internal/domain"

type FamilyListRepository interface {
	CreateBulkFamilyList(familyLists []domain.FamilyList) error
	CreateFamilyList(familyList *domain.FamilyList) error
	GetFamilyListByID(flId uint) (*domain.FamilyList, error)
	GetAllFamilyListsByUserID(cstId uint) ([]domain.FamilyList, error)
	UpdateFamilyList(familyList *domain.FamilyList) error
	DeleteFamilyList(flId uint) error
}

type FamilyListService interface {
	CreateBulkFamilyList(familyLists []domain.FamilyList) error
	CreateFamilyList(familyList *domain.FamilyList) error
	GetFamilyListByID(flId uint) (*domain.FamilyList, error)
	GetAllFamilyListsByUserID(cstId uint) ([]domain.FamilyList, error)
	UpdateFamilyList(familyList *domain.FamilyList) error
	DeleteFamilyList(flId uint) error
}
