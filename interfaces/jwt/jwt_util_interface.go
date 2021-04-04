package jwt

import (
	"authentication/models"
	"github.com/dgrijalva/jwt-go"
)

type JwtUtil interface {
	CreateToken(user *models.User) *string
	ValidateToken(token string) (*jwt.Token, *error)
}
