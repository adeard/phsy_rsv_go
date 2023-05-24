package city

import "phsy_rsv_go/domain"

type Service interface {
	FindAll() ([]domain.City, error)
	Delete(ID int) (domain.City, error)
	FindByID(ID int) (domain.City, error)
	Create(cityrequest domain.CityRequest) (domain.City, error)
	Update(ID int, cityrequest domain.CityRequest) (domain.City, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]domain.City, error) {
	cities, err := s.repository.FindAll()

	return cities, err
}

func (s *service) FindByID(ID int) (domain.City, error) {
	city, err := s.repository.FindByID(ID)

	return city, err
}

func (s *service) Create(cityrequest domain.CityRequest) (domain.City, error) {

	newCity := domain.City{
		Name:       cityrequest.Name,
		IsActive:   cityrequest.IsActive,
		ProvinceId: cityrequest.ProvinceId,
	}

	city, err := s.repository.Create(newCity)

	return city, err
}

func (s *service) Update(ID int, cityrequest domain.CityRequest) (domain.City, error) {

	city, err := s.repository.FindByID(ID)
	if err != nil {
		return domain.City{}, err
	}

	city.Name = cityrequest.Name
	city.IsActive = cityrequest.IsActive
	city.ProvinceId = cityrequest.ProvinceId

	newCity, err := s.repository.Update(city)

	return newCity, err
}

func (s *service) Delete(ID int) (domain.City, error) {

	city, err := s.repository.FindByID(ID)
	if err != nil {
		return domain.City{}, err
	}

	city, err = s.repository.Delete(city)
	if err != nil {
		return domain.City{}, err
	}

	return city, err
}
