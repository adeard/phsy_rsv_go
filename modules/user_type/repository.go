package usertype

import (
	"phsy_rsv_go/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.UserType, error)
	FindByID(ID int) (domain.UserType, error)
	Delete(usertype domain.UserType) (domain.UserType, error)
	Create(usertype domain.UserType) (domain.UserType, error)
	Update(usertype domain.UserType) (domain.UserType, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.UserType, error) {
	var usertypes []domain.UserType
	err := r.db.Find(&usertypes).Error

	return usertypes, err
}

func (r *repository) FindByID(ID int) (domain.UserType, error) {
	var usertype domain.UserType
	err := r.db.Where("id = ?", ID).First(&usertype).Error

	return usertype, err
}

func (r *repository) Create(usertype domain.UserType) (domain.UserType, error) {
	err := r.db.Create(&usertype).Error

	return usertype, err
}

func (r *repository) Update(usertype domain.UserType) (domain.UserType, error) {
	err := r.db.Debug().Save(&usertype).Error

	return usertype, err
}

func (r *repository) Delete(usertype domain.UserType) (domain.UserType, error) {
	err := r.db.Debug().Delete(&usertype).Error

	return usertype, err
}
