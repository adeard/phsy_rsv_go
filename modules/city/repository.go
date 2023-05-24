package city

import (
	"phsy_rsv_go/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.City, error)
	FindByID(ID int) (domain.City, error)
	Delete(city domain.City) (domain.City, error)
	Create(city domain.City) (domain.City, error)
	Update(city domain.City) (domain.City, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.City, error) {
	var cities []domain.City
	err := r.db.Find(&cities).Error

	return cities, err
}

func (r *repository) FindByID(ID int) (domain.City, error) {
	var city domain.City
	err := r.db.Preload("Province").Where("id = ?", ID).First(&city).Error

	return city, err
}

func (r *repository) Create(city domain.City) (domain.City, error) {
	err := r.db.Create(&city).Error

	return city, err
}

func (r *repository) Update(city domain.City) (domain.City, error) {
	err := r.db.Debug().Save(&city).Error

	return city, err
}

func (r *repository) Delete(city domain.City) (domain.City, error) {
	err := r.db.Debug().Delete(&city).Error

	return city, err
}
