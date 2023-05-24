package rate

import (
	"phsy_rsv_go/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.Rate, error)
	FindByID(ID int) (domain.Rate, error)
	Delete(rate domain.Rate) (domain.Rate, error)
	Create(rate domain.Rate) (domain.Rate, error)
	Update(rate domain.Rate) (domain.Rate, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.Rate, error) {
	var rates []domain.Rate
	err := r.db.Find(&rates).Error

	return rates, err
}

func (r *repository) FindByID(ID int) (domain.Rate, error) {
	var rate domain.Rate
	err := r.db.Where("id = ?", ID).First(&rate).Error

	return rate, err
}

func (r *repository) Create(rate domain.Rate) (domain.Rate, error) {
	err := r.db.Create(&rate).Error

	return rate, err
}

func (r *repository) Update(rate domain.Rate) (domain.Rate, error) {
	err := r.db.Debug().Save(&rate).Error

	return rate, err
}

func (r *repository) Delete(rate domain.Rate) (domain.Rate, error) {
	err := r.db.Debug().Delete(&rate).Error

	return rate, err
}
