package rate

import "phsy_rsv_go/domain"

type Service interface {
	FindAll() ([]domain.Rate, error)
	Delete(ID int) (domain.Rate, error)
	FindByID(ID int) (domain.Rate, error)
	Create(raterequest domain.RateRequest) (domain.Rate, error)
	Update(ID int, raterequest domain.RateRequest) (domain.Rate, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]domain.Rate, error) {
	rates, err := s.repository.FindAll()

	return rates, err
}

func (s *service) FindByID(ID int) (domain.Rate, error) {
	rate, err := s.repository.FindByID(ID)

	return rate, err
}

func (s *service) Create(raterequest domain.RateRequest) (domain.Rate, error) {

	newRate := domain.Rate{
		UserId: raterequest.UserId,
		Rates:  raterequest.Rates,
	}

	rate, err := s.repository.Create(newRate)

	return rate, err
}

func (s *service) Update(ID int, raterequest domain.RateRequest) (domain.Rate, error) {

	rate, err := s.repository.FindByID(ID)
	if err != nil {
		return domain.Rate{}, err
	}

	rate.UserId = raterequest.UserId
	rate.Rates = raterequest.Rates

	newrate, err := s.repository.Update(rate)

	return newrate, err
}

func (s *service) Delete(ID int) (domain.Rate, error) {

	rate, err := s.repository.FindByID(ID)
	if err != nil {
		return domain.Rate{}, err
	}

	rate, err = s.repository.Delete(rate)
	if err != nil {
		return domain.Rate{}, err
	}

	return rate, err
}
