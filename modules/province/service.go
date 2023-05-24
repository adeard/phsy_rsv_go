package province

import "phsy_rsv_go/domain"

type Service interface {
	FindAll() ([]domain.Province, error)
	Delete(ID int) (domain.Province, error)
	FindByID(ID int) (domain.Province, error)
	Create(provincerequest domain.ProvinceRequest) (domain.Province, error)
	Update(ID int, provincerequest domain.ProvinceRequest) (domain.Province, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]domain.Province, error) {
	provinces, err := s.repository.FindAll()

	return provinces, err
}

func (s *service) FindByID(ID int) (domain.Province, error) {
	province, err := s.repository.FindByID(ID)

	return province, err
}

func (s *service) Create(provincerequest domain.ProvinceRequest) (domain.Province, error) {

	newProvince := domain.Province{
		Name:     provincerequest.Name,
		IsActive: provincerequest.IsActive,
	}

	province, err := s.repository.Create(newProvince)

	return province, err
}

func (s *service) Update(ID int, provincerequest domain.ProvinceRequest) (domain.Province, error) {

	province, err := s.repository.FindByID(ID)
	if err != nil {
		return domain.Province{}, err
	}

	province.Name = provincerequest.Name
	province.IsActive = provincerequest.IsActive

	newprovince, err := s.repository.Update(province)

	return newprovince, err
}

func (s *service) Delete(ID int) (domain.Province, error) {

	province, err := s.repository.FindByID(ID)
	if err != nil {
		return domain.Province{}, err
	}

	province, err = s.repository.Delete(province)
	if err != nil {
		return domain.Province{}, err
	}

	return province, err
}
