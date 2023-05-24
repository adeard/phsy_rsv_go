package province

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Province, error)
	FindByID(ID int) (Province, error)
	Delete(province Province) (Province, error)
	Create(province Province) (Province, error)
	Update(province Province) (Province, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Province, error) {
	var provinces []Province
	err := r.db.Find(&provinces).Error

	return provinces, err
}

func (r *repository) FindByID(ID int) (Province, error) {
	var province Province
	err := r.db.Where("id = ?", ID).First(&province).Error

	return province, err
}

func (r *repository) Create(province Province) (Province, error) {
	err := r.db.Create(&province).Error

	return province, err
}

func (r *repository) Update(province Province) (Province, error) {
	err := r.db.Debug().Save(&province).Error

	return province, err
}

func (r *repository) Delete(province Province) (Province, error) {
	err := r.db.Debug().Delete(&province).Error

	return province, err
}
