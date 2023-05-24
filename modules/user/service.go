package user

import (
	"html"
	"phsy_rsv_go/domain"
	"phsy_rsv_go/utils"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	FindAll() ([]domain.User, error)
	Delete(ID int) (domain.User, error)
	FindByID(ID int) (domain.User, error)
	Login(loginrequest domain.LoginRequest) (string, error)
	Create(registerrequest domain.RegisterRequest) (domain.User, error)
	Update(ID int, updaterequest domain.UpdateRequest) (domain.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]domain.User, error) {
	users, err := s.repository.FindAll()

	return users, err
}

func (s *service) FindByID(ID int) (domain.User, error) {
	user, err := s.repository.FindByID(ID)

	return user, err
}

func (s *service) Create(registerrequest domain.RegisterRequest) (domain.User, error) {

	newUser, _ := hashedUser(domain.User{
		Username: registerrequest.Username,
		Password: registerrequest.Password,
	})

	user, err := s.repository.Create(newUser)

	return user, err
}

func (s *service) Login(loginrequest domain.LoginRequest) (string, error) {

	userCheck, err := s.repository.FindByUsername(loginrequest.Username)
	if err != nil {
		return "", err
	}

	err = verifyPassword(loginrequest.Password, userCheck.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(userCheck.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) Update(ID int, updaterequest domain.UpdateRequest) (domain.User, error) {

	user, err := s.repository.FindByID(ID)
	if err != nil {
		return domain.User{}, err
	}

	user.Address = updaterequest.Address
	user.BirthDate = updaterequest.BirthDate
	user.Email = updaterequest.Email
	user.Gender = updaterequest.Gender
	user.ImgProfile = updaterequest.ImgProfile
	user.IsActive = updaterequest.IsActive
	user.Name = updaterequest.Name

	newuser, err := s.repository.Update(user)

	return newuser, err
}

func (s *service) Delete(ID int) (domain.User, error) {

	user, err := s.repository.FindByID(ID)
	if err != nil {
		return domain.User{}, err
	}

	user, err = s.repository.Delete(user)
	if err != nil {
		return domain.User{}, err
	}

	return user, err
}

func hashedUser(u domain.User) (domain.User, error) {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return u, err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return u, nil

}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
