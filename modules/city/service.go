package city

type Service interface {
	FindAll() ([]City, error)
	Delete(ID int) (City, error)
	FindByID(ID int) (City, error)
	Create(cityrequest CityRequest) (City, error)
	Update(ID int, cityrequest CityRequest) (City, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]City, error) {
	cities, err := s.repository.FindAll()

	return cities, err
}

func (s *service) FindByID(ID int) (City, error) {
	city, err := s.repository.FindByID(ID)

	return city, err
}

func (s *service) Create(cityrequest CityRequest) (City, error) {

	newCity := City{
		Name:       cityrequest.Name,
		IsActive:   cityrequest.IsActive,
		ProvinceId: cityrequest.ProvinceId,
	}

	city, err := s.repository.Create(newCity)

	return city, err
}

func (s *service) Update(ID int, cityrequest CityRequest) (City, error) {

	city, err := s.repository.FindByID(ID)
	if err != nil {
		return City{}, err
	}

	city.Name = cityrequest.Name
	city.IsActive = cityrequest.IsActive
	city.ProvinceId = cityrequest.ProvinceId

	newCity, err := s.repository.Update(city)

	return newCity, err
}

func (s *service) Delete(ID int) (City, error) {

	city, err := s.repository.FindByID(ID)
	if err != nil {
		return City{}, err
	}

	city, err = s.repository.Delete(city)
	if err != nil {
		return City{}, err
	}

	return city, err
}
