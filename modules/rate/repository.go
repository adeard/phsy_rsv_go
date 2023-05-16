package rate

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Rate, error)
	FindByID(ID int) (Rate, error)
	Delete(rate Rate) (Rate, error)
	Create(rate Rate) (Rate, error)
	Update(rate Rate) (Rate, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Rate, error) {
	var rates []Rate
	err := r.db.Find(&rates).Error

	return rates, err
}

func (r *repository) FindByID(ID int) (Rate, error) {
	var rate Rate
	err := r.db.Where("id = ?", ID).First(&rate).Error

	return rate, err
}

func (r *repository) Create(rate Rate) (Rate, error) {
	err := r.db.Create(&rate).Error

	return rate, err
}

func (r *repository) Update(rate Rate) (Rate, error) {
	err := r.db.Debug().Save(&rate).Error

	return rate, err
}

func (r *repository) Delete(rate Rate) (Rate, error) {
	err := r.db.Debug().Delete(&rate).Error

	return rate, err
}
