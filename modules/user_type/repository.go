package usertype

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]UserType, error)
	FindByID(ID int) (UserType, error)
	Delete(usertype UserType) (UserType, error)
	Create(usertype UserType) (UserType, error)
	Update(usertype UserType) (UserType, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]UserType, error) {
	var usertypes []UserType
	err := r.db.Find(&usertypes).Error

	return usertypes, err
}

func (r *repository) FindByID(ID int) (UserType, error) {
	var usertype UserType
	err := r.db.Where("id = ?", ID).First(&usertype).Error

	return usertype, err
}

func (r *repository) Create(usertype UserType) (UserType, error) {
	err := r.db.Create(&usertype).Error

	return usertype, err
}

func (r *repository) Update(usertype UserType) (UserType, error) {
	err := r.db.Debug().Save(&usertype).Error

	return usertype, err
}

func (r *repository) Delete(usertype UserType) (UserType, error) {
	err := r.db.Debug().Delete(&usertype).Error

	return usertype, err
}
