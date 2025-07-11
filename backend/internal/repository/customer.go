package repository

import (
	"github.com/hnifmaghfur/bookingtogo/internal/domain"
	"github.com/hnifmaghfur/bookingtogo/internal/interfaces"
	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

var _ interfaces.CustomerRepository = (*customerRepository)(nil)

func NewCustomerRepository(db *gorm.DB) interfaces.CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) CreateCustomer(customer *domain.Customer) error {
	return r.db.Save(customer).Error
}

func (r *customerRepository) GetCustomerByID(id uint) (*domain.Customer, error) {
	var customer domain.Customer
	// Preload kedua relasi
	err := r.db.Preload("FamilyList").Preload("Nationality").First(&customer, id).Error
	return &customer, err
}

func (r *customerRepository) GetAllCustomers() ([]domain.Customer, error) {
	var customers []domain.Customer
	err := r.db.Preload("FamilyList").Preload("Nationality").Find(&customers).Error
	return customers, err
}

func (r *customerRepository) UpdateCustomer(customer *domain.Customer) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	existingCustomer := &domain.Customer{}
	if err := tx.Preload("FamilyList").First(existingCustomer, customer.CstID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Hapus FamilyList lama yang terkait dengan customer ini
	if err := tx.Where("cst_id = ?", customer.CstID).Delete(&domain.FamilyList{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update data customer utama. Pastikan NationalityID diupdate juga.
	// Jika customer.NationalityID adalah nil, itu akan mengeset kolom nationality_id di DB menjadi NULL.
	if err := tx.Session(&gorm.Session{FullSaveAssociations: false}).Updates(customer).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Tambahkan FamilyList baru
	for i := range customer.FamilyList {
		customer.FamilyList[i].CstID = customer.CstID
		if err := tx.Create(&customer.FamilyList[i]).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *customerRepository) DeleteCustomer(id uint) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Hapus FamilyList terkait (jika ON DELETE CASCADE tidak diatur di DB)
	if err := tx.Where("cst_id = ?", id).Delete(&domain.FamilyList{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Hapus Customer. Karena nationality_id adalah FK di Customer dan tidak ada ON DELETE CASCADE
	// ke Nationality, Nationality yang direferensikan tidak akan terhapus.
	// Kolom nationality_id di Customer akan menjadi NULL atau customer dihapus.
	if err := tx.Delete(&domain.Customer{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
