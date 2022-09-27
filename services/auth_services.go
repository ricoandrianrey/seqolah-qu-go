package services

import (
	"seqolah-qu/entities/dto"
	"seqolah-qu/entities/models"
	"seqolah-qu/repositories"
)

type AuthService interface {
	Register()
	Login()
}

type AuthServiceImpl struct {
	UserRepository *repositories.UserRepositoryImpl
}

func NewAuthService(userRepo *repositories.UserRepositoryImpl) *AuthServiceImpl {
	return &AuthServiceImpl{userRepo}
}

func (s *AuthServiceImpl) Register(data dto.RegisterDto) interface{} {
	existingUser, _ := s.UserRepository.FindOne(models.User{Email: data.Email, Role: data.Role})

	if existingUser.Email != "" {
		return nil
	}

	newUser, _ := s.UserRepository.Create(models.User{Email: data.Email, Role: data.Role, Password: data.Password, Status: "active"})
	return newUser
}

func (s *AuthServiceImpl) Login(data dto.LoginDto) interface{} {
	// login logic
	return nil
}
