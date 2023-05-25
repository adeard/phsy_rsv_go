package userlevel

import "phsy_rsv_go/domain"

type Service interface {
	FindAll() ([]domain.UserLevel, error)
	Delete(ID int) (domain.UserLevel, error)
	FindByID(ID int) (domain.UserLevel, error)
	Create(userlevelrequest domain.UserLevelRequest) (domain.UserLevel, error)
	Update(ID int, userlevelrequest domain.UserLevelRequest) (domain.UserLevel, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]domain.UserLevel, error) {
	userlevels, err := s.repository.FindAll()

	return userlevels, err
}

func (s *service) FindByID(ID int) (domain.UserLevel, error) {
	userlevel, err := s.repository.FindByID(ID)

	return userlevel, err
}

func (s *service) Create(userlevelrequest domain.UserLevelRequest) (domain.UserLevel, error) {

	newUserLevel := domain.UserLevel{
		Name:     userlevelrequest.Name,
		IsActive: userlevelrequest.IsActive,
	}

	userlevel, err := s.repository.Create(newUserLevel)

	return userlevel, err
}

func (s *service) Update(ID int, userlevelrequest domain.UserLevelRequest) (domain.UserLevel, error) {

	userlevel, err := s.repository.FindByID(ID)
	if err != nil {
		return domain.UserLevel{}, err
	}

	userlevel.Name = userlevelrequest.Name
	userlevel.IsActive = userlevelrequest.IsActive

	newuserlevel, err := s.repository.Update(userlevel)

	return newuserlevel, err
}

func (s *service) Delete(ID int) (domain.UserLevel, error) {

	userlevel, err := s.repository.FindByID(ID)
	if err != nil {
		return domain.UserLevel{}, err
	}

	userlevel, err = s.repository.Delete(userlevel)
	if err != nil {
		return domain.UserLevel{}, err
	}

	return userlevel, err
}
