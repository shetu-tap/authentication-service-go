package authentication

import "authentication/models"

type AuthenticationRepository interface {
	CreateUser(user models.User) *models.User
	FindUserByUsername(username string) (*models.User, *error)
}
