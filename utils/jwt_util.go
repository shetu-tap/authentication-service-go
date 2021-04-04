package utils

import (
	jwt_go "github.com/dgrijalva/jwt-go"
	"authentication/interfaces/jwt"
	"authentication/models"
	"time"
)

type jwtUtil struct {

}

func (jwtUtil *jwtUtil) ValidateToken(token string) (*jwt_go.Token, *error) {
	t, err := jwt_go.ParseWithClaims(token, &jwt_go.StandardClaims{}, func(token *jwt_go.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return nil, &err
	}
	return t, nil
}

var (
	SECRET_KEY = "this-is-a-secret-key"
)

func (jwtUtil *jwtUtil) CreateToken(user *models.User) *string {
	claims := &jwt_go.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   user.Username,
	}
	token := jwt_go.NewWithClaims(jwt_go.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return nil
	}
	return &signedToken
}

func GetJwtUtil() jwt.JwtUtil{
	return &jwtUtil{}
}