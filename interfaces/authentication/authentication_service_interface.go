package authentication

import (
	authentication_dto "authentication/dto/authentication"
	"authentication/models"
)

type AuthenticationService interface {
	CreateUser(body authentication_dto.RegistrationBody) (*models.User, *error)
	FindUserByUsername(username string) (*models.User, *error)
}
