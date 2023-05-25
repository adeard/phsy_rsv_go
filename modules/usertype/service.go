package usertype

import "phsy_rsv_go/domain"

type Service interface {
	FindAll() ([]domain.UserType, error)
	Delete(ID int) (domain.UserType, error)
	FindByID(ID int) (domain.UserType, error)
	Create(usertyperequest domain.UserTypeRequest) (domain.UserType, error)
	Update(ID int, usertyperequest domain.UserTypeRequest) (domain.UserType, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]domain.UserType, error) {
	usertypes, err := s.repository.FindAll()

	return usertypes, err
}

func (s *service) FindByID(ID int) (domain.UserType, error) {
	usertype, err := s.repository.FindByID(ID)

	return usertype, err
}

func (s *service) Create(usertyperequest domain.UserTypeRequest) (domain.UserType, error) {

	newUserType := domain.UserType{
		Name:     usertyperequest.Name,
		IsActive: usertyperequest.IsActive,
	}

	usertype, err := s.repository.Create(newUserType)

	return usertype, err
}

func (s *service) Update(ID int, usertyperequest domain.UserTypeRequest) (domain.UserType, error) {

	usertype, err := s.repository.FindByID(ID)
	if err != nil {
		return domain.UserType{}, err
	}

	usertype.Name = usertyperequest.Name
	usertype.IsActive = usertyperequest.IsActive

	newuser, err := s.repository.Update(usertype)

	return newuser, err
}

func (s *service) Delete(ID int) (domain.UserType, error) {

	usertype, err := s.repository.FindByID(ID)
	if err != nil {
		return domain.UserType{}, err
	}

	usertype, err = s.repository.Delete(usertype)
	if err != nil {
		return domain.UserType{}, err
	}

	return usertype, err
}
