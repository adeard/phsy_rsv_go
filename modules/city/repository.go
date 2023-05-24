package city

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]City, error)
	FindByID(ID int) (City, error)
	Delete(city City) (City, error)
	Create(city City) (City, error)
	Update(city City) (City, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]City, error) {
	var cities []City
	err := r.db.Find(&cities).Error

	return cities, err
}

func (r *repository) FindByID(ID int) (City, error) {
	var city City
	err := r.db.Where("id = ?", ID).First(&city).Error

	return city, err
}

func (r *repository) Create(city City) (City, error) {
	err := r.db.Create(&city).Error

	return city, err
}

func (r *repository) Update(city City) (City, error) {
	err := r.db.Debug().Save(&city).Error

	return city, err
}

func (r *repository) Delete(city City) (City, error) {
	err := r.db.Debug().Delete(&city).Error

	return city, err
}
