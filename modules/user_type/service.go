package usertype

type Service interface {
	FindAll() ([]UserType, error)
	Delete(ID int) (UserType, error)
	FindByID(ID int) (UserType, error)
	Create(usertyperequest UserTypeRequest) (UserType, error)
	Update(ID int, usertyperequest UserTypeRequest) (UserType, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]UserType, error) {
	usertypes, err := s.repository.FindAll()

	return usertypes, err
}

func (s *service) FindByID(ID int) (UserType, error) {
	usertype, err := s.repository.FindByID(ID)

	return usertype, err
}

func (s *service) Create(usertyperequest UserTypeRequest) (UserType, error) {

	newUserType := UserType{
		Name:     usertyperequest.Name,
		IsActive: usertyperequest.IsActive,
	}

	usertype, err := s.repository.Create(newUserType)

	return usertype, err
}

func (s *service) Update(ID int, usertyperequest UserTypeRequest) (UserType, error) {

	usertype, err := s.repository.FindByID(ID)
	if err != nil {
		return UserType{}, err
	}

	usertype.Name = usertyperequest.Name
	usertype.IsActive = usertyperequest.IsActive

	newuser, err := s.repository.Update(usertype)

	return newuser, err
}

func (s *service) Delete(ID int) (UserType, error) {

	usertype, err := s.repository.FindByID(ID)
	if err != nil {
		return UserType{}, err
	}

	usertype, err = s.repository.Delete(usertype)
	if err != nil {
		return UserType{}, err
	}

	return usertype, err
}
