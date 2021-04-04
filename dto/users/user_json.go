package users

import "github.com/google/uuid"

type UserJson struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
}
