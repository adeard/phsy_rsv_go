package userlevel

import (
	"phsy_rsv_go/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.UserLevel, error)
	FindByID(ID int) (domain.UserLevel, error)
	Delete(userlevel domain.UserLevel) (domain.UserLevel, error)
	Create(userlevel domain.UserLevel) (domain.UserLevel, error)
	Update(userlevel domain.UserLevel) (domain.UserLevel, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.UserLevel, error) {
	var userlevels []domain.UserLevel
	err := r.db.Find(&userlevels).Error

	return userlevels, err
}

func (r *repository) FindByID(ID int) (domain.UserLevel, error) {
	var userlevel domain.UserLevel
	err := r.db.Where("id = ?", ID).First(&userlevel).Error

	return userlevel, err
}

func (r *repository) Create(userlevel domain.UserLevel) (domain.UserLevel, error) {
	err := r.db.Create(&userlevel).Error

	return userlevel, err
}

func (r *repository) Update(userlevel domain.UserLevel) (domain.UserLevel, error) {
	err := r.db.Debug().Save(&userlevel).Error

	return userlevel, err
}

func (r *repository) Delete(userlevel domain.UserLevel) (domain.UserLevel, error) {
	err := r.db.Debug().Delete(&userlevel).Error

	return userlevel, err
}
