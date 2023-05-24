package user

import (
	"phsy_rsv_go/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.User, error)
	FindByID(ID int) (domain.User, error)
	FindByUsername(username string) (domain.User, error)
	Delete(user domain.User) (domain.User, error)
	Create(user domain.User) (domain.User, error)
	Update(user domain.User) (domain.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.
		Find(&users).Error

	return users, err
}

func (r *repository) FindByID(ID int) (domain.User, error) {
	var user domain.User
	err := r.db.
		Preload("UserType").
		Preload("UserLevel").
		Where("id = ?", ID).
		First(&user).Error

	return user, err
}

func (r *repository) FindByUsername(username string) (domain.User, error) {
	var user domain.User
	err := r.db.
		Where("username = ?", username).
		First(&user).Error

	return user, err
}

func (r *repository) Create(user domain.User) (domain.User, error) {
	err := r.db.
		Preload("UserType").
		Create(&user).Error

	return user, err
}

func (r *repository) Update(user domain.User) (domain.User, error) {
	err := r.db.
		Save(&user).Error

	return user, err
}

func (r *repository) Delete(user domain.User) (domain.User, error) {
	err := r.db.
		Debug().
		Delete(&user).Error

	return user, err
}
