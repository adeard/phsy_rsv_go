package userlevel

type Service interface {
	FindAll() ([]UserLevel, error)
	Delete(ID int) (UserLevel, error)
	FindByID(ID int) (UserLevel, error)
	Create(userlevelrequest UserLevelRequest) (UserLevel, error)
	Update(ID int, userlevelrequest UserLevelRequest) (UserLevel, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]UserLevel, error) {
	userlevels, err := s.repository.FindAll()

	return userlevels, err
}

func (s *service) FindByID(ID int) (UserLevel, error) {
	userlevel, err := s.repository.FindByID(ID)

	return userlevel, err
}

func (s *service) Create(userlevelrequest UserLevelRequest) (UserLevel, error) {

	newUserLevel := UserLevel{
		Name:     userlevelrequest.Name,
		IsActive: userlevelrequest.IsActive,
	}

	userlevel, err := s.repository.Create(newUserLevel)

	return userlevel, err
}

func (s *service) Update(ID int, userlevelrequest UserLevelRequest) (UserLevel, error) {

	userlevel, err := s.repository.FindByID(ID)
	if err != nil {
		return UserLevel{}, err
	}

	userlevel.Name = userlevelrequest.Name
	userlevel.IsActive = userlevelrequest.IsActive

	newuserlevel, err := s.repository.Update(userlevel)

	return newuserlevel, err
}

func (s *service) Delete(ID int) (UserLevel, error) {

	userlevel, err := s.repository.FindByID(ID)
	if err != nil {
		return UserLevel{}, err
	}

	userlevel, err = s.repository.Delete(userlevel)
	if err != nil {
		return UserLevel{}, err
	}

	return userlevel, err
}
