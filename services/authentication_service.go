package services

import (
	authentication_dto "authentication/dto/authentication"
	"authentication/interfaces/authentication"
	"authentication/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	MIN_COST = 4
	MAX_COST = 15
	DEFAULT_COST = 10
)

type authenticationService struct {
	authenticationRepository authentication.AuthenticationRepository
}

func (authenticationService *authenticationService) FindUserByUsername(username string) (*models.User, *error) {
	user, err := authenticationService.authenticationRepository.FindUserByUsername(username)
	return user, err
}

func (authenticationService *authenticationService) CreateUser(body authentication_dto.RegistrationBody) (*models.User, *error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), DEFAULT_COST)
	if err != nil {
		return nil, &err
	}
	user := models.User{
		Name:     body.Name,
		Username: body.Username,
		Password: string(hashedPassword),
	}
	return authenticationService.authenticationRepository.CreateUser(user), nil
}

func GetAuthenticationService(repository authentication.AuthenticationRepository) authentication.AuthenticationService{
	return &authenticationService{
		authenticationRepository: repository,
	}
}