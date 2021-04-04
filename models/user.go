package models

import (
	"authentication/dto/users"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id uuid.UUID `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	user.Id = uuid.New()
	return
}

func (user User) ToUserJSON() users.UserJson {
	return users.UserJson{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
	}
}

func (user User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}