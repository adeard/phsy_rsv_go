package rate

type Service interface {
	FindAll() ([]Rate, error)
	Delete(ID int) (Rate, error)
	FindByID(ID int) (Rate, error)
	Create(raterequest RateRequest) (Rate, error)
	Update(ID int, raterequest RateRequest) (Rate, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Rate, error) {
	rates, err := s.repository.FindAll()

	return rates, err
}

func (s *service) FindByID(ID int) (Rate, error) {
	rate, err := s.repository.FindByID(ID)

	return rate, err
}

func (s *service) Create(raterequest RateRequest) (Rate, error) {

	newRate := Rate{
		UserId: raterequest.UserId,
		Rates:  raterequest.Rates,
	}

	rate, err := s.repository.Create(newRate)

	return rate, err
}

func (s *service) Update(ID int, raterequest RateRequest) (Rate, error) {

	rate, err := s.repository.FindByID(ID)
	if err != nil {
		return Rate{}, err
	}

	rate.UserId = raterequest.UserId
	rate.Rates = raterequest.Rates

	newrate, err := s.repository.Update(rate)

	return newrate, err
}

func (s *service) Delete(ID int) (Rate, error) {

	rate, err := s.repository.FindByID(ID)
	if err != nil {
		return Rate{}, err
	}

	rate, err = s.repository.Delete(rate)
	if err != nil {
		return Rate{}, err
	}

	return rate, err
}
