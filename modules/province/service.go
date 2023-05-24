package province

type Service interface {
	FindAll() ([]Province, error)
	Delete(ID int) (Province, error)
	FindByID(ID int) (Province, error)
	Create(provincerequest ProvinceRequest) (Province, error)
	Update(ID int, provincerequest ProvinceRequest) (Province, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Province, error) {
	provinces, err := s.repository.FindAll()

	return provinces, err
}

func (s *service) FindByID(ID int) (Province, error) {
	province, err := s.repository.FindByID(ID)

	return province, err
}

func (s *service) Create(provincerequest ProvinceRequest) (Province, error) {

	newProvince := Province{
		Name:     provincerequest.Name,
		IsActive: provincerequest.IsActive,
	}

	province, err := s.repository.Create(newProvince)

	return province, err
}

func (s *service) Update(ID int, provincerequest ProvinceRequest) (Province, error) {

	province, err := s.repository.FindByID(ID)
	if err != nil {
		return Province{}, err
	}

	province.Name = provincerequest.Name
	province.IsActive = provincerequest.IsActive

	newprovince, err := s.repository.Update(province)

	return newprovince, err
}

func (s *service) Delete(ID int) (Province, error) {

	province, err := s.repository.FindByID(ID)
	if err != nil {
		return Province{}, err
	}

	province, err = s.repository.Delete(province)
	if err != nil {
		return Province{}, err
	}

	return province, err
}
