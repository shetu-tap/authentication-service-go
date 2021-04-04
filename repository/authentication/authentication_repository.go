package authentication

import (
	"authentication/db"
	authentication_interfaces "authentication/interfaces/authentication"
	"authentication/models"
	"errors"
	"gorm.io/gorm"
)

type authenticationRepository struct {
	db *gorm.DB
}

func (authenticationRepository *authenticationRepository) FindUserByUsername(username string) (*models.User, *error) {
	var user models.User
	result := authenticationRepository.db.Where("username = ?", username).First(&user)

	if errors.Is(result.Error,gorm.ErrRecordNotFound) {
		return nil, &result.Error
	}
	return &user, nil
}

func (authenticationRepository *authenticationRepository) CreateUser(user models.User) *models.User {
	authenticationRepository.db.Create(&user)
	return &user
}

func GetAuthenticationRepository() authentication_interfaces.AuthenticationRepository {
	postgresDB := db.SetupPostgres()
	return &authenticationRepository{
		db: postgresDB,
	}
}