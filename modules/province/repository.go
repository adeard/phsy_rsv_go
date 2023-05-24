package province

import (
	"fmt"
	"phsy_rsv_go/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.Province, error)
	FindByID(ID int) (domain.Province, error)
	Delete(province domain.Province) (domain.Province, error)
	Create(province domain.Province) (domain.Province, error)
	Update(province domain.Province) (domain.Province, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.Province, error) {
	var provinces []domain.Province
	err := r.db.Find(&provinces).Error

	return provinces, err
}

func (r *repository) FindByID(ID int) (domain.Province, error) {
	var province domain.Province
	err := r.db.Where("id = ?", ID).First(&province).Error

	fmt.Println(province)

	return province, err
}

func (r *repository) Create(province domain.Province) (domain.Province, error) {
	err := r.db.Create(&province).Error

	return province, err
}

func (r *repository) Update(province domain.Province) (domain.Province, error) {
	err := r.db.Debug().Save(&province).Error

	return province, err
}

func (r *repository) Delete(province domain.Province) (domain.Province, error) {
	err := r.db.Debug().Delete(&province).Error

	return province, err
}
