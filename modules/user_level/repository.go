package userlevel

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]UserLevel, error)
	FindByID(ID int) (UserLevel, error)
	Delete(userlevel UserLevel) (UserLevel, error)
	Create(userlevel UserLevel) (UserLevel, error)
	Update(userlevel UserLevel) (UserLevel, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]UserLevel, error) {
	var userlevels []UserLevel
	err := r.db.Find(&userlevels).Error

	return userlevels, err
}

func (r *repository) FindByID(ID int) (UserLevel, error) {
	var userlevel UserLevel
	err := r.db.Where("id = ?", ID).First(&userlevel).Error

	return userlevel, err
}

func (r *repository) Create(userlevel UserLevel) (UserLevel, error) {
	err := r.db.Create(&userlevel).Error

	return userlevel, err
}

func (r *repository) Update(userlevel UserLevel) (UserLevel, error) {
	err := r.db.Debug().Save(&userlevel).Error

	return userlevel, err
}

func (r *repository) Delete(userlevel UserLevel) (UserLevel, error) {
	err := r.db.Debug().Delete(&userlevel).Error

	return userlevel, err
}
