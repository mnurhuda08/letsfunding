package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	CheckEmailAvailabilty(input CheckEmailAvailabiltyInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserByID(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(PasswordHash)
	user.Role = "user"

	newUser, err := s.repository.Insert(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	pasword := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("User Not Found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(pasword))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) CheckEmailAvailabilty(input CheckEmailAvailabiltyInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return true, nil
	}

	return false, nil

}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	user.AvatarFileName = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil

}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("No User Found")
	}

	return user, nil
}
