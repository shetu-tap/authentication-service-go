package users

import (
	"authentication/dto/errors"
	"authentication/interfaces/authentication"
	"authentication/interfaces/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userController struct {
	authentictionService authentication.AuthenticationService
}

func (userController *userController) Me(c *gin.Context) {
	username, err := c.Get("username")
	if !err {
		c.JSON(http.StatusForbidden, errors.ErrorResponse{
			Message: "Forbidden",
			Code:    http.StatusForbidden,
		})
		return
	}

	user, e := userController.authentictionService.FindUserByUsername(username.(string))
	if e != nil {
		c.JSON(http.StatusNotFound, errors.ErrorResponse{
			Message: "User not found",
			Code:    http.StatusNotFound,
		})
		return
	}
	c.JSON(http.StatusOK, user.ToUserJSON())
}

func GetUserController(service authentication.AuthenticationService) users.UserController{
	return &userController{
		authentictionService: service,
	}
}
